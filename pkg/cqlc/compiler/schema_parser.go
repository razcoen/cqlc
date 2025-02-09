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
	schemaBuilder    *schemaBuilder
	keyspaceBuilders map[string]*keyspaceBuilder
	err              error
}

func newSchemaParserTreeListener() *schemaParserTreeListener {
	l := &schemaParserTreeListener{
		BaseCQLParserListener: &antlrcql.BaseCQLParserListener{},
		schemaBuilder:         newSchemaBuilder(),
	}
	return l
}

func (l *schemaParserTreeListener) recordError(err error) {
	l.err = errors.Join(l.err, err)
}

func (l *schemaParserTreeListener) EnterAlterTable(ctx *antlrcql.AlterTableContext) {
	var keyspaceContext *antlrcql.KeyspaceContext
	var tableContext *antlrcql.TableContext
	var alterTableOperationContext *antlrcql.AlterTableOperationContext
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.KeyspaceContext:
			keyspaceContext = child
		case *antlrcql.TableContext:
			tableContext = child
		case *antlrcql.AlterTableOperationContext:
			alterTableOperationContext = child
		}
	}

	// Evaluate the table that is being altered.
	var keyspace string
	var tableName string
	if keyspaceContext != nil {
		keyspace = keyspaceContext.GetText()
	}
	if tableContext != nil {
		tableName = tableContext.GetText()
	}
	if tableName == "" {
		l.recordError(fmt.Errorf("alter table: failed to find the table name in alter table context"))
		return
	}
	var kb *keyspaceBuilder
	var foundKeyspace bool
	for k, b := range l.keyspaceBuilders {
		if k == keyspace {
			kb = b
			foundKeyspace = true
			break
		}
	}
	if !foundKeyspace {
		l.recordError(fmt.Errorf("alter table: did not find keyspace %q within the schema context", keyspace))
		return
	}
	var table *sdk.Table
	for _, t := range kb.tables {
		if t.Name == tableName {
			table = t
		}
	}
	if table == nil {
		l.recordError(fmt.Errorf("alter table: did not find table %q in keyspace %q within the schema context", tableName, keyspace))
		return
	}

	for _, child := range alterTableOperationContext.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.AlterTableAddContext:
			alterTableAddContext := child
			for _, child := range alterTableAddContext.GetChildren() {
				switch child := child.(type) {
				case *antlrcql.AlterTableColumnDefinitionContext:
					alterTableColumnDefinitionContext := child
					col, _, err := parseColumnDefinitionContext(alterTableColumnDefinitionContext)
					if err != nil {
						l.recordError(fmt.Errorf("alter table: parse alter table column definition context: %w", err))
						continue
					}
					columnAlreadyExists := false
					for _, c := range table.Columns {
						if c.Name == col.Name {
							columnAlreadyExists = true
							break
						}
					}
					if columnAlreadyExists {
						// TODO: Logger
						continue
					}
					table.Columns = append(table.Columns, col)
				}
			}
		case *antlrcql.AlterTableDropColumnsContext:
			droppedColumnNames := map[string]bool{}
			alterTableDropColumnsContext := child
			for _, child := range alterTableDropColumnsContext.GetChildren() {
				switch child := child.(type) {
				case *antlrcql.AlterTableDropColumnListContext:
					alterTableDropColumnListContext := child
					for _, child := range alterTableDropColumnListContext.GetChildren() {
						switch child := child.(type) {
						case *antlrcql.ColumnContext:
							columnName := child.GetText()
							droppedColumnNames[columnName] = true
						}
					}
				}
			}
			var columns []*sdk.Column
			for _, col := range table.Columns {
				if dropped := droppedColumnNames[col.Name]; !dropped {
					columns = append(columns, col)
				}
			}
			table.Columns = columns

		default:
		}
	}
}

func (l *schemaParserTreeListener) EnterDropTable(ctx *antlrcql.DropTableContext) {
	var keyspaceContext *antlrcql.KeyspaceContext
	var tableContext *antlrcql.TableContext
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.KeyspaceContext:
			keyspaceContext = child
		case *antlrcql.TableContext:
			tableContext = child
		}
	}

	// Evaluate the table that is being altered.
	var keyspace string
	var tableName string
	if keyspaceContext != nil {
		keyspace = keyspaceContext.GetText()
	}
	if tableContext != nil {
		tableName = tableContext.GetText()
	}
	if tableName == "" {
		l.recordError(fmt.Errorf("drop table: failed to find the table name in alter table context"))
		return
	}
	var kb *keyspaceBuilder
	var foundKeyspace bool
	for k, b := range l.keyspaceBuilders {
		if k == keyspace {
			kb = b
			foundKeyspace = true
			break
		}
	}
	if !foundKeyspace {
		return
	}
	var tables []*sdk.Table
	for _, t := range kb.tables {
		if t.Name == tableName {
			continue
		}
		tables = append(tables, t)
	}
	// TODO: The builder abstraction fails here.
	kb.tables = tables
	delete(kb.tableNames, tableName)
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
			column, isPrimaryKey, err := parseColumnDefinitionContext(child)
			if err != nil {
				l.recordError(fmt.Errorf("create table: parse column definition context: %w", err))
				continue
			}
			tableBuilder = tableBuilder.withColumn(column.Name, column.DataType)
			if isPrimaryKey {
				tableBuilder = tableBuilder.withPrimaryKey(column.Name)
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

type columnDefinitionContext interface {
	*antlrcql.ColumnDefinitionContext | *antlrcql.AlterTableColumnDefinitionContext
	GetChildren() []antlr.Tree
	GetText() string
}

func parseColumnDefinitionContext[T columnDefinitionContext](ctx T) (col *sdk.Column, isPrimaryKey bool, err error) {
	var columnName string
	var columnType string
	for _, child := range ctx.GetChildren() {
		switch child := child.(type) {
		case *antlrcql.ColumnContext:
			columnName = child.GetText()
		case *antlrcql.DataTypeContext:
			columnType = child.GetText()
		case *antlrcql.PrimaryKeyColumnContext:
			isPrimaryKey = true
		}
	}
	if columnName == "" {
		return nil, false, fmt.Errorf("did not find column name on this column definition context: %q", ctx.GetText())
	}
	if columnType == "" {
		return nil, false, fmt.Errorf("did not find column type on this column definition context: %q", ctx.GetText())
	}
	dataType := gocqlhelpers.ParseCassandraType(columnType,
		log.New(os.Stdout, "", 0), // TODO: Propragate or handle this logger better
	)
	column := &sdk.Column{Name: columnName, DataType: dataType}
	return column, isPrimaryKey, nil
}
