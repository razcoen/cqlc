package cql

import (
	"errors"
	"fmt"
	"maps"
	"slices"
)

const defaultKeyspaceName = "system"

type SchemaBuilder struct {
	keyspaces map[string]*Keyspace
	err       error
}

func NewSchemaBuilder() *SchemaBuilder {
	return &SchemaBuilder{keyspaces: make(map[string]*Keyspace)}
}

func (sb *SchemaBuilder) WithKeyspace(keyspace *Keyspace) *SchemaBuilder {
	if _, ok := sb.keyspaces[keyspace.Name]; ok {
		sb.err = errors.Join(sb.err, fmt.Errorf(`keyspace "%s" already exists`, keyspace.Name))
		return sb
	}
	sb.keyspaces[keyspace.Name] = keyspace
	return sb
}

func (sb *SchemaBuilder) Build() (*Schema, error) {
	if _, ok := sb.keyspaces[defaultKeyspaceName]; !ok {
		defaultKeyspace, err := NewDefaultKeyspaceBuilder().Build()
		if err != nil {
			return nil, fmt.Errorf("build default keyspace: %w", err)
		}
		sb.keyspaces[defaultKeyspace.Name] = defaultKeyspace
	}
	if sb.err != nil {
		return nil, fmt.Errorf("error during schema build process: %w", sb.err)
	}
	return &Schema{Keyspaces: slices.Collect(maps.Values(sb.keyspaces))}, nil
}

type KeyspaceBuilder struct {
	name   string
	tables map[string]*Table
	err    error
}

func NewDefaultKeyspaceBuilder() *KeyspaceBuilder {
	return NewKeyspaceBuilder(defaultKeyspaceName)
}

func NewKeyspaceBuilder(name string) *KeyspaceBuilder {
	return &KeyspaceBuilder{
		name:   name,
		tables: make(map[string]*Table),
	}
}

func (kb *KeyspaceBuilder) WithTable(table *Table) *KeyspaceBuilder {
	if _, ok := kb.tables[table.Name]; ok {
		kb.err = errors.Join(kb.err, fmt.Errorf(`table "%s" already exists`, table.Name))
		return kb
	}
	kb.tables[table.Name] = table
	return kb
}

func (kb *KeyspaceBuilder) Build() (*Keyspace, error) {
	if kb.err != nil {
		return nil, fmt.Errorf("error during keyspace build process: %w", kb.err)
	}
	return &Keyspace{
		Name:   kb.name,
		Tables: slices.Collect(maps.Values(kb.tables)),
	}, nil
}

type TableBuilder struct {
	name    string
	columns map[string]*DataType
	err     error
}

func NewTableBuilder(name string) *TableBuilder {
	return &TableBuilder{
		name:    name,
		columns: make(map[string]*DataType),
	}
}
func (tb *TableBuilder) WithColumn(columnName string, columnType *DataType) *TableBuilder {
	if _, ok := tb.columns[columnName]; ok {
		tb.err = errors.Join(tb.err, fmt.Errorf(`column "%s" already exists`, columnName))
		return tb
	}
	tb.columns[columnName] = columnType
	return tb
}

func (tb *TableBuilder) Build() (*Table, error) {
	if tb.err != nil {
		return nil, fmt.Errorf("error during table build process: %w", tb.err)
	}
	columns := make([]*Column, 0, len(tb.columns))
	for c, t := range tb.columns {
		columns = append(columns, &Column{Name: c, DataType: t})
	}
	return &Table{Name: tb.name, Columns: columns}, nil
}
