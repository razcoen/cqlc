package sdk

import (
	"testing"

	"github.com/razcoen/cqlc/pkg/cqlc/gocqlhelpers"
	"github.com/stretchr/testify/assert"
)

func TestSchemaString(t *testing.T) {
	schema := &Schema{
		Keyspaces: []*Keyspace{{
			Name: "ks1",
			Tables: []*Table{
				{
					Name: "tbl1",
					Columns: []*Column{
						{Name: "col1", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "col2", DataType: gocqlhelpers.NewTypeTime()},
						{Name: "col3", DataType: gocqlhelpers.NewTypeInt()}},
					PrimaryKey: &PrimaryKey{
						PartitionKey:  []string{"col1"},
						ClusteringKey: []string{"col3"},
					}},
			}}}}
	expected := `{
  "keyspaces": [
    {
      "name": "ks1",
      "tables": [
        {
          "name": "tbl1",
          "columns": [
            {
              "data_type": "uuid",
              "name": "col1"
            },
            {
              "data_type": "time",
              "name": "col2"
            },
            {
              "data_type": "int",
              "name": "col3"
            }
          ],
          "primary_key": {
            "partition_key": [
              "col1"
            ],
            "clustering_key": [
              "col3"
            ]
          }
        }
      ]
    }
  ]
}`
	assert.Equal(t, expected, schema.String())
}

func TestKeyspaceString(t *testing.T) {
	keyspace := &Keyspace{
		Name: "ks1",
		Tables: []*Table{
			{
				Name: "tbl1",
				Columns: []*Column{
					{Name: "col1", DataType: gocqlhelpers.NewTypeUUID()},
					{Name: "col2", DataType: gocqlhelpers.NewTypeTime()},
					{Name: "col3", DataType: gocqlhelpers.NewTypeInt()}},
				PrimaryKey: &PrimaryKey{
					PartitionKey:  []string{"col1"},
					ClusteringKey: []string{"col3"},
				}},
		}}
	expected := `{
  "name": "ks1",
  "tables": [
    {
      "name": "tbl1",
      "columns": [
        {
          "data_type": "uuid",
          "name": "col1"
        },
        {
          "data_type": "time",
          "name": "col2"
        },
        {
          "data_type": "int",
          "name": "col3"
        }
      ],
      "primary_key": {
        "partition_key": [
          "col1"
        ],
        "clustering_key": [
          "col3"
        ]
      }
    }
  ]
}`
	assert.Equal(t, expected, keyspace.String())
}

func TestTableString(t *testing.T) {
	table := &Table{
		Name: "tbl1",
		Columns: []*Column{
			{Name: "col1", DataType: gocqlhelpers.NewTypeUUID()},
			{Name: "col2", DataType: gocqlhelpers.NewTypeTime()},
			{Name: "col3", DataType: gocqlhelpers.NewTypeInt()}},
		PrimaryKey: &PrimaryKey{
			PartitionKey:  []string{"col1"},
			ClusteringKey: []string{"col3"},
		},
	}
	expected := `{
  "name": "tbl1",
  "columns": [
    {
      "data_type": "uuid",
      "name": "col1"
    },
    {
      "data_type": "time",
      "name": "col2"
    },
    {
      "data_type": "int",
      "name": "col3"
    }
  ],
  "primary_key": {
    "partition_key": [
      "col1"
    ],
    "clustering_key": [
      "col3"
    ]
  }
}`
	assert.Equal(t, expected, table.String())
}

func TestColumnString(t *testing.T) {
	column := &Column{Name: "col1", DataType: gocqlhelpers.NewTypeUUID()}
	expected := `{
  "data_type": "uuid",
  "name": "col1"
}`
	assert.Equal(t, expected, column.String())
	column = &Column{Name: "col2", DataType: gocqlhelpers.NewTypeCustom("hello")}
	expected = `{
  "data_type": "custom:hello",
  "name": "col2"
}`
	assert.Equal(t, expected, column.String())
}
