package cqlc

import (
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"maps"
	"slices"
	"strings"
)

const defaultKeyspaceName = ""

type SchemaBuilder struct {
	keyspaceNames map[string]bool
	keyspaces     []*Keyspace
	err           error
}

func NewSchemaBuilder() *SchemaBuilder {
	return &SchemaBuilder{keyspaceNames: make(map[string]bool)}
}

func (sb *SchemaBuilder) WithKeyspace(keyspace *Keyspace) *SchemaBuilder {
	if _, ok := sb.keyspaceNames[keyspace.Name]; ok {
		sb.err = errors.Join(sb.err, fmt.Errorf(`keyspace "%s" already exists`, keyspace.Name))
		return sb
	}
	sb.keyspaceNames[keyspace.Name] = true
	sb.keyspaces = append(sb.keyspaces, keyspace)
	return sb
}

func (sb *SchemaBuilder) Build() (*Schema, error) {
	if _, ok := sb.keyspaceNames[defaultKeyspaceName]; !ok {
		defaultKeyspace, err := NewDefaultKeyspaceBuilder().Build()
		if err != nil {
			return nil, fmt.Errorf("build default keyspace: %w", err)
		}
		sb.keyspaceNames[defaultKeyspace.Name] = true
		sb.keyspaces = append(sb.keyspaces, defaultKeyspace)
	}
	if sb.err != nil {
		return nil, fmt.Errorf("error during schema build process: %w", sb.err)
	}
	return &Schema{Keyspaces: sb.keyspaces}, nil
}

type KeyspaceBuilder struct {
	name       string
	tableNames map[string]bool
	tables     []*Table
	err        error
}

func NewDefaultKeyspaceBuilder() *KeyspaceBuilder {
	return NewKeyspaceBuilder(defaultKeyspaceName)
}

func NewKeyspaceBuilder(name string) *KeyspaceBuilder {
	return &KeyspaceBuilder{
		name:       name,
		tableNames: make(map[string]bool),
	}
}

func (kb *KeyspaceBuilder) WithTable(table *Table) *KeyspaceBuilder {
	if _, ok := kb.tableNames[table.Name]; ok {
		kb.err = errors.Join(kb.err, fmt.Errorf(`table "%s" already exists`, table.Name))
		return kb
	}
	kb.tableNames[table.Name] = true
	kb.tables = append(kb.tables, table)
	return kb
}

func (kb *KeyspaceBuilder) Build() (*Keyspace, error) {
	if kb.err != nil {
		return nil, fmt.Errorf("error during keyspace build process: %w", kb.err)
	}
	return &Keyspace{Name: kb.name, Tables: kb.tables}, nil
}

type TableBuilder struct {
	name          string
	columnNames   map[string]bool
	partitionKey  []string
	clusteringKey []string
	columns       []*Column
	err           error
}

func NewTableBuilder(name string) *TableBuilder {
	return &TableBuilder{
		name:        name,
		columnNames: make(map[string]bool),
	}
}

func (tb *TableBuilder) WithPrimaryKey(columnName string) *TableBuilder {
	tb.partitionKey = append(tb.partitionKey, columnName)
	return tb
}

func (tb *TableBuilder) WithPartitionKey(columnName string) *TableBuilder {
	tb.partitionKey = append(tb.partitionKey, columnName)
	return tb
}

func (tb *TableBuilder) WithClusteringKey(columnName string) *TableBuilder {
	tb.clusteringKey = append(tb.clusteringKey, columnName)
	return tb
}

func (tb *TableBuilder) WithColumn(columnName string, columnType gocql.TypeInfo) *TableBuilder {
	if _, ok := tb.columnNames[columnName]; ok {
		tb.err = errors.Join(tb.err, fmt.Errorf(`column "%s" already exists`, columnName))
		return tb
	}
	tb.columnNames[columnName] = true
	tb.columns = append(tb.columns, &Column{Name: columnName, DataType: columnType})
	return tb
}

func (tb *TableBuilder) Build() (*Table, error) {
	if tb.err != nil {
		return nil, fmt.Errorf("error during table build process: %w", tb.err)
	}
	// Order columns according to cassandra natural order:
	// 1. Primary key by columns order
	// 2. Other columns sort alphabetically
	columnByName := make(map[string]*Column, len(tb.columns))
	for _, column := range tb.columns {
		columnByName[column.Name] = column
	}
	var columns []*Column
	for _, column := range tb.partitionKey {
		columns = append(columns, columnByName[column])
		delete(columnByName, column)
	}
	for _, column := range tb.clusteringKey {
		columns = append(columns, columnByName[column])
		delete(columnByName, column)
	}
	columnsLeft := slices.Collect(maps.Values(columnByName))
	slices.SortFunc(columnsLeft, func(a, b *Column) int { return strings.Compare(a.Name, b.Name) })
	columns = append(columns, columnsLeft...)
	return &Table{Name: tb.name, Columns: columns, PrimaryKey: &PrimaryKey{PartitionKey: tb.partitionKey, ClusteringKey: tb.clusteringKey}}, nil
}
