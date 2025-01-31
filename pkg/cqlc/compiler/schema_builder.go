package compiler

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
)

const defaultKeyspaceName = ""

type schemaBuilder struct {
	keyspaceNames map[string]bool
	keyspaces     []*sdk.Keyspace
	err           error
}

func newSchemaBuilder() *schemaBuilder {
	return &schemaBuilder{keyspaceNames: make(map[string]bool)}
}

func (sb *schemaBuilder) withKeyspace(keyspace *sdk.Keyspace) *schemaBuilder {
	if _, ok := sb.keyspaceNames[keyspace.Name]; ok {
		sb.err = errors.Join(sb.err, fmt.Errorf(`keyspace "%s" already exists`, keyspace.Name))
		return sb
	}
	sb.keyspaceNames[keyspace.Name] = true
	sb.keyspaces = append(sb.keyspaces, keyspace)
	return sb
}

func (sb *schemaBuilder) build() (*sdk.Schema, error) {
	if _, ok := sb.keyspaceNames[defaultKeyspaceName]; !ok {
		defaultKeyspace, err := newDefaultKeyspaceBuilder().build()
		if err != nil {
			return nil, fmt.Errorf("build default keyspace: %w", err)
		}
		sb.keyspaceNames[defaultKeyspace.Name] = true
		sb.keyspaces = append(sb.keyspaces, defaultKeyspace)
	}
	if sb.err != nil {
		return nil, fmt.Errorf("error during schema build process: %w", sb.err)
	}
	return &sdk.Schema{Keyspaces: sb.keyspaces}, nil
}

type keyspaceBuilder struct {
	name       string
	tableNames map[string]bool
	tables     []*sdk.Table
	err        error
}

func newDefaultKeyspaceBuilder() *keyspaceBuilder {
	return newKeyspaceBuilder(defaultKeyspaceName)
}

func newKeyspaceBuilder(name string) *keyspaceBuilder {
	return &keyspaceBuilder{
		name:       name,
		tableNames: make(map[string]bool),
	}
}

func (kb *keyspaceBuilder) withTable(table *sdk.Table) *keyspaceBuilder {
	if _, ok := kb.tableNames[table.Name]; ok {
		kb.err = errors.Join(kb.err, fmt.Errorf(`table "%s" already exists`, table.Name))
		return kb
	}
	kb.tableNames[table.Name] = true
	kb.tables = append(kb.tables, table)
	return kb
}

func (kb *keyspaceBuilder) build() (*sdk.Keyspace, error) {
	if kb.err != nil {
		return nil, fmt.Errorf("error during keyspace build process: %w", kb.err)
	}
	return &sdk.Keyspace{Name: kb.name, Tables: kb.tables}, nil
}

type tableBuilder struct {
	name          string
	columnNames   map[string]bool
	partitionKey  []string
	clusteringKey []string
	columns       []*sdk.Column
	err           error
}

func newTableBuilder(name string) *tableBuilder {
	return &tableBuilder{
		name:        name,
		columnNames: make(map[string]bool),
	}
}

func (tb *tableBuilder) withPrimaryKey(columnName string) *tableBuilder {
	tb.partitionKey = append(tb.partitionKey, columnName)
	return tb
}

func (tb *tableBuilder) withPartitionKey(columnName string) *tableBuilder {
	tb.partitionKey = append(tb.partitionKey, columnName)
	return tb
}

func (tb *tableBuilder) withClusteringKey(columnName string) *tableBuilder {
	tb.clusteringKey = append(tb.clusteringKey, columnName)
	return tb
}

func (tb *tableBuilder) withColumn(columnName string, columnType gocql.TypeInfo) *tableBuilder {
	if _, ok := tb.columnNames[columnName]; ok {
		tb.err = errors.Join(tb.err, fmt.Errorf(`column "%s" already exists`, columnName))
		return tb
	}
	tb.columnNames[columnName] = true
	tb.columns = append(tb.columns, &sdk.Column{Name: columnName, DataType: columnType})
	return tb
}

func (tb *tableBuilder) build() (*sdk.Table, error) {
	if tb.err != nil {
		return nil, fmt.Errorf("error during table build process: %w", tb.err)
	}
	// Order columns according to cassandra natural order:
	// 1. Primary key by columns order
	// 2. Other columns sort alphabetically
	columnByName := make(map[string]*sdk.Column, len(tb.columns))
	for _, column := range tb.columns {
		columnByName[column.Name] = column
	}
	var columns []*sdk.Column
	for _, column := range tb.partitionKey {
		columns = append(columns, columnByName[column])
		delete(columnByName, column)
	}
	for _, column := range tb.clusteringKey {
		columns = append(columns, columnByName[column])
		delete(columnByName, column)
	}
	columnsLeft := slices.Collect(maps.Values(columnByName))
	slices.SortFunc(columnsLeft, func(a, b *sdk.Column) int { return strings.Compare(a.Name, b.Name) })
	columns = append(columns, columnsLeft...)
	return &sdk.Table{Name: tb.name, Columns: columns, PrimaryKey: &sdk.PrimaryKey{PartitionKey: tb.partitionKey, ClusteringKey: tb.clusteringKey}}, nil
}
