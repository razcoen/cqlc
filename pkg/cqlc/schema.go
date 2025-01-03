package cqlc

import (
	"github.com/gocql/gocql"
	"strings"
)

type Schema struct {
	Keyspaces []*Keyspace
}

func (s *Schema) String() string {
	sb := strings.Builder{}
	sb.WriteString("{keyspaces:[")
	for i, k := range s.Keyspaces {
		sb.WriteString(k.String())
		if i < len(s.Keyspaces)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]}")
	return sb.String()
}

type Keyspace struct {
	Name   string
	Tables []*Table
}

func (k *Keyspace) String() string {
	sb := strings.Builder{}
	sb.WriteString("{name:")
	sb.WriteString(k.Name)
	sb.WriteString(",tables:[")
	for i, t := range k.Tables {
		sb.WriteString(t.String())
		if i < len(k.Tables)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]}")
	return sb.String()
}

type Table struct {
	Name    string
	Columns []*Column
}

func (t *Table) String() string {
	sb := strings.Builder{}
	sb.WriteString("{name:")
	sb.WriteString(t.Name)
	sb.WriteString(",columns:[")
	for i, c := range t.Columns {
		sb.WriteString(c.String())
		if i < len(t.Columns)-1 {
			sb.WriteString(",")
		}
	}
	sb.WriteString("]}")
	return sb.String()
}

type Column struct {
	Name     string
	DataType gocql.TypeInfo
}

func (c *Column) String() string {
	sb := strings.Builder{}
	sb.WriteString("{name:")
	sb.WriteString(c.Name)
	sb.WriteString(",type:")
	sb.WriteString(c.DataType.Type().String())
	if c.DataType.Type() == gocql.TypeCustom {
		sb.WriteString(":")
		sb.WriteString(c.DataType.Custom())
	}
	sb.WriteString("}")
	return sb.String()
}
