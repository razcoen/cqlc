package cql

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/razcoen/cqlc/pkg/cql/parser"
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
		lexer := parser.NewCQLLexer(antlr.NewInputStream(stmt))
		lexer.RemoveErrorListeners()
		lexer.AddErrorListener(el)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewCQLParser(stream)
		p.RemoveErrorListeners()
		p.AddErrorListener(el)
		antlr.ParseTreeWalkerDefault.Walk(l, p.Cql())
	}
	if el.errors != nil {
		return nil, errors.Join(el.errors...)
	}
	if l.err != nil {
		return nil, fmt.Errorf("error during traversal: %w", l.err)
	}
	defaultKeyspace, err := l.defaultKeyspaceBuilder.Build()
	if err != nil {
		return nil, fmt.Errorf("build default keyspace: %w", err)
	}
	return l.schemaBuilder.WithKeyspace(defaultKeyspace).Build()
}


type errorListener struct {
	*antlr.DefaultErrorListener
	errors []error
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

func newErrorListener() *errorListener {
	return &errorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

type schemaParserTreeListener struct {
	*parser.BaseCQLParserListener
	schemaBuilder          *SchemaBuilder
	defaultKeyspaceBuilder *KeyspaceBuilder
	err                    error
}

func newSchemaParserTreeListener() *schemaParserTreeListener {
	l := &schemaParserTreeListener{
		BaseCQLParserListener: &parser.BaseCQLParserListener{},
		// TODO: Do we even need to support more than one keyspace as part of the code generation?
		schemaBuilder:          NewSchemaBuilder(),
		defaultKeyspaceBuilder: NewDefaultKeyspaceBuilder(),
	}
	return l
}

func (l *schemaParserTreeListener) EnterCreateTable(ctx *parser.CreateTableContext) {
	var columnDefinitionListContext *parser.ColumnDefinitionListContext
	for _, child := range ctx.GetChildren() {
		ctx, ok := child.(*parser.ColumnDefinitionListContext)
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
		columnDefinitionContext, ok := child.(*parser.ColumnDefinitionContext)
		if !ok {
			continue
		}
		var columnName string
		var columnType string
		for _, child := range columnDefinitionContext.GetChildren() {
			switch child := child.(type) {
			case *parser.ColumnContext:
				columnName = child.GetText()
			case *parser.DataTypeContext:
				columnType = child.GetText()
			default:
			}
		}
		dataType, err := ParseDataType(columnType)
		if err != nil {
			l.err = errors.Join(l.err, fmt.Errorf("parse data type: %w", err))
			continue
		}
		tableBuilder = tableBuilder.WithColumn(columnName, dataType)
	}
	table, err := tableBuilder.Build()
	if err != nil {
		l.err = errors.Join(l.err, fmt.Errorf("build table: %w", err))
		return
	}

	// TODO: Extend support to many keyspaces
	l.defaultKeyspaceBuilder = l.defaultKeyspaceBuilder.WithTable(table)
}
