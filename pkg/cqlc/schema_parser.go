package cqlc

import (
	"errors"
	"fmt"
	"github.com/razcoen/cqlc/pkg/antlrhelpers"
	"github.com/razcoen/cqlc/pkg/gocqlhelpers"
	"log"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/internal/antlrcql"
)

type SchemaParser struct{}

func NewSchemaParser() *SchemaParser {
	return &SchemaParser{}
}

func (sp *SchemaParser) Parse(cql string) (*Schema, error) {
	l := newSchemaParserTreeListener()
	el := newErrorListener()
	for _, stmt := range strings.Split(cql, ";") {
		stmt := strings.TrimSpace(stmt)
		lexer := antlrcql.NewCQLLexer(antlr.NewInputStream(stmt))
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(el)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := antlrcql.NewCQLParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(el)
		t := p.Cql()
		antlrhelpers.PrintTree(t)
		antlr.ParseTreeWalkerDefault.Walk(l, t)
	}
	if el.errors != nil {
		return nil, errors.Join(el.errors...)
	}
	if l.err != nil {
		return nil, fmt.Errorf("error during traversal: %w", l.err)
	}
	for _, kb := range l.keyspaceBuilders {
		ks, err := kb.Build()
		if err != nil {
			return nil, fmt.Errorf("build keyspace: %w", err)
		}
		l.schemaBuilder = l.schemaBuilder.WithKeyspace(ks)
	}
	return l.schemaBuilder.Build()
}

type errorListener struct {
	*antlr.DefaultErrorListener
	errors []error
}

func newErrorListener() *errorListener {
	return &errorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	token, ok := offendingSymbol.(antlr.Token)
	var tokenText string
	if ok {
		tokenText = token.GetText()
	}
	// TODO: The following are ignored but requires fix through the g4 syntax files.
	if tokenText == "<EOF>" || tokenText == "<" {
		return
	}
	err := fmt.Errorf(`syntax error "%s" in line %d:%d %s`, tokenText, line, column, msg)
	l.errors = append(l.errors, err)
}

type schemaParserTreeListener struct {
	*antlrcql.BaseCQLParserListener
	schemaBuilder          *SchemaBuilder
	keyspaceBuilders       map[string]*KeyspaceBuilder
	defaultKeyspaceBuilder *KeyspaceBuilder
	err                    error
}

func newSchemaParserTreeListener() *schemaParserTreeListener {
	l := &schemaParserTreeListener{
		BaseCQLParserListener: &antlrcql.BaseCQLParserListener{},
		// TODO: Do we even need to support more than one keyspace as part of the code generation?
		schemaBuilder:          NewSchemaBuilder(),
		defaultKeyspaceBuilder: NewDefaultKeyspaceBuilder(),
	}
	return l
}

func (l *schemaParserTreeListener) EnterCreateTable(ctx *antlrcql.CreateTableContext) {
	// TODO: Refactor to more resilient
	var columnDefinitionListContext *antlrcql.ColumnDefinitionListContext
	for _, child := range ctx.GetChildren() {
		ctx, ok := child.(*antlrcql.ColumnDefinitionListContext)
		if !ok {
			continue
		}
		columnDefinitionListContext = ctx
		break
	}
	if columnDefinitionListContext == nil {
		return
	}

	tableBuilder := NewTableBuilder(ctx.Table().GetText())
	for _, child := range columnDefinitionListContext.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.PrimaryKeyElementContext:
			if child.GetChildCount() < 4 {
				continue
			}
			primaryKeyDefinitionContext, ok := child.GetChild(3).(*antlrcql.PrimaryKeyDefinitionContext)
			if !ok {
				continue
			}
			handlePartitionKeyContext := func(partitionKeyContext *antlrcql.PartitionKeyContext) {
				for _, c := range partitionKeyContext.GetChildren() {
					columnContext, ok := c.(*antlrcql.ColumnContext)
					if !ok {
						continue
					}
					tableBuilder = tableBuilder.WithPartitionKey(columnContext.GetText())
				}
			}
			keyContext := primaryKeyDefinitionContext.GetChild(0) // Either CompositeKeyContext or CompoundKeyContext
			partitionKeyContext, ok := keyContext.GetChild(0).(*antlrcql.PartitionKeyContext)
			if ok {
				handlePartitionKeyContext(partitionKeyContext)
			}
			partitionKeyListContext, ok := keyContext.GetChild(1).(*antlrcql.PartitionKeyListContext)
			if ok {
				for _, c := range partitionKeyListContext.GetChildren() {
					partitionKeyContext, ok := c.(*antlrcql.PartitionKeyContext)
					if ok {
						handlePartitionKeyContext(partitionKeyContext)
					}
				}
			}
			if keyContext.GetChildCount() < 3 {
				continue
			}
			clusteringKeyContext, ok := keyContext.GetChild(2).(*antlrcql.ClusteringKeyListContext)
			if !ok {
				clusteringKeyContext, ok = keyContext.GetChild(4).(*antlrcql.ClusteringKeyListContext)
				if !ok {
					continue
				}
			}
			for _, c := range clusteringKeyContext.GetChildren() {
				clusteringKeyContext, ok := c.(*antlrcql.ClusteringKeyContext)
				if !ok {
					continue
				}
				for _, col := range clusteringKeyContext.GetChildren() {
					columnContext, ok := col.(*antlrcql.ColumnContext)
					if !ok {
						continue
					}
					tableBuilder = tableBuilder.WithClusteringKey(columnContext.GetText())

				}
			}
		case *antlrcql.ColumnDefinitionContext:
			var columnName string
			var columnType string
			isPrimaryKey := false
			for _, child := range child.GetChildren() {
				switch child := child.(type) {
				case *antlrcql.ColumnContext:
					columnName = child.GetText()
				case *antlrcql.DataTypeContext:
					columnType = child.GetText()
				case *antlrcql.PrimaryKeyColumnContext:
					isPrimaryKey = true
				default:
				}
			}
			// TODO: Logger? Errors?
			ti := gocqlhelpers.ParseCassandraType(columnType, log.New(os.Stdout, "", 0))
			tableBuilder = tableBuilder.WithColumn(columnName, ti)
			if isPrimaryKey {
				tableBuilder = tableBuilder.WithPrimaryKey(columnName)
			}
		}
	}
	table, err := tableBuilder.Build()
	if err != nil {
		l.err = errors.Join(l.err, fmt.Errorf("build table: %w", err))
		return
	}

	keyspace := defaultKeyspaceName
	if ctx.Keyspace() != nil && ctx.Keyspace().GetText() != "" {
		keyspace = ctx.Keyspace().GetText()
	}
	if l.keyspaceBuilders == nil {
		l.keyspaceBuilders = make(map[string]*KeyspaceBuilder, 1)
	}
	kb, ok := l.keyspaceBuilders[keyspace]
	if !ok {
		kb = NewKeyspaceBuilder(keyspace)
	}
	kb = kb.WithTable(table)
	l.keyspaceBuilders[keyspace] = kb
}
