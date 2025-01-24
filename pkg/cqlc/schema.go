package cqlc

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
)

type Schema struct {
	Keyspaces []*Keyspace `json:"keyspaces"`
}

func (s *Schema) String() string {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err).Error()
	}
	return string(b)
}

type Keyspace struct {
	Name   string   `json:"name"`
	Tables []*Table `json:"tables"`
}

func (k *Keyspace) String() string {
	b, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err).Error()
	}
	return string(b)
}

type Table struct {
	Name       string      `json:"name"`
	Columns    []*Column   `json:"columns"`
	PrimaryKey *PrimaryKey `json:"primary_key"`
}

type PrimaryKey struct {
	PartitionKey  []string `json:"partition_key"`
	ClusteringKey []string `json:"clustering_key"`
}

func (t *Table) String() string {
	b, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err).Error()
	}
	return string(b)
}

type Column struct {
	Name     string         // JSON marshalling is customer, therefore no tags
	DataType gocql.TypeInfo // JSON marshalling is customer, therefore no tags
}

func (c *Column) MarshalJSON() ([]byte, error) {
	node := map[string]string{"name": c.Name, "data_type": c.DataType.Type().String()}
	if c.DataType.Type() == gocql.TypeCustom {
		node["data_type"] = node["data_type"] + ":" + c.DataType.Custom()
	}
	return json.Marshal(node)
}

func (c *Column) String() string {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err).Error()
	}
	return string(b)
}
