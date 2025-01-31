package compiler

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/compiler/internal/antlrcql"
	"github.com/razcoen/cqlc/pkg/cqlc/gocqlhelpers"
)

type SchemaParser struct{}

func NewSchemaParser() *SchemaParser {
	return &SchemaParser{}
}

func (sp *SchemaParser) Parse(cql string) (*sdk.Schema, error) {
	cqls, err := antlrParse(cql)
	if err != nil {
		return nil, fmt.Errorf("antlr parse cql: %w", err)
	}

	l := newSchemaParserTreeListener()
	for _, cql := range cqls {
		antlr.ParseTreeWalkerDefault.Walk(l, cql)
	}
	if l.err != nil {
		return nil, fmt.Errorf("error during traversal: %w", l.err)
	}
	for _, kb := range l.keyspaceBuilders {
		ks, err := kb.build()
		if err != nil {
			return nil, fmt.Errorf("build keyspace: %w", err)
		}
		l.schemaBuilder = l.schemaBuilder.withKeyspace(ks)
	}
	return l.schemaBuilder.build()
}

type schemaParserTreeListener struct {
	*antlrcql.BaseCQLParserListener
	schemaBuilder          *schemaBuilder
	keyspaceBuilders       map[string]*keyspaceBuilder
	defaultKeyspaceBuilder *keyspaceBuilder
	err                    error
}

func newSchemaParserTreeListener() *schemaParserTreeListener {
	l := &schemaParserTreeListener{
		BaseCQLParserListener: &antlrcql.BaseCQLParserListener{},
		// TODO: Do we even need to support more than one keyspace as part of the code generation?
		schemaBuilder:          newSchemaBuilder(),
		defaultKeyspaceBuilder: newDefaultKeyspaceBuilder(),
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

	tableBuilder := newTableBuilder(ctx.Table().GetText())
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
					tableBuilder = tableBuilder.withPartitionKey(columnContext.GetText())
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
					tableBuilder = tableBuilder.withClusteringKey(columnContext.GetText())

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
			tableBuilder = tableBuilder.withColumn(columnName, ti)
			if isPrimaryKey {
				tableBuilder = tableBuilder.withPrimaryKey(columnName)
			}
		}
	}
	table, err := tableBuilder.build()
	if err != nil {
		l.err = errors.Join(l.err, fmt.Errorf("build table: %w", err))
		return
	}

	keyspace := defaultKeyspaceName
	if ctx.Keyspace() != nil && ctx.Keyspace().GetText() != "" {
		keyspace = ctx.Keyspace().GetText()
	}
	if l.keyspaceBuilders == nil {
		l.keyspaceBuilders = make(map[string]*keyspaceBuilder, 1)
	}
	kb, ok := l.keyspaceBuilders[keyspace]
	if !ok {
		kb = newKeyspaceBuilder(keyspace)
	}
	kb = kb.withTable(table)
	l.keyspaceBuilders[keyspace] = kb
}
