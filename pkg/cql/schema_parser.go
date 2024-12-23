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
	for _, stmt := range strings.Split(cql, ";") {
		lexer := parser.NewCqlLexer(antlr.NewInputStream(stmt))
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewCqlParser(stream)
		p.RemoveErrorListeners()
		antlr.ParseTreeWalkerDefault.Walk(l, p.Cql())
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

type schemaParserTreeListener struct {
	*parser.BaseCqlParserListener
	schemaBuilder          *SchemaBuilder
	defaultKeyspaceBuilder *KeyspaceBuilder
	err                    error
}

func newSchemaParserTreeListener() *schemaParserTreeListener {
	l := &schemaParserTreeListener{
		BaseCqlParserListener: &parser.BaseCqlParserListener{},
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
