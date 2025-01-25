package compiler

import (
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/gocqlhelpers"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaBuilder(t *testing.T) {
	t1, err := NewTableBuilder("t1").
		WithColumn("c1", gocqlhelpers.NewTypeText()).
		WithColumn("c2", gocqlhelpers.NewTypeSet(gocqlhelpers.NewTypeTimestamp())).
		WithPrimaryKey("c1").
		Build()
	require.NoError(t, err)
	k1, err := NewKeyspaceBuilder("k1").
		WithTable(t1).
		Build()
	require.NoError(t, err)
	s1, err := NewSchemaBuilder().WithKeyspace(k1).Build()
	require.NoError(t, err)
	expected := &sdk.Schema{
		Keyspaces: []*sdk.Keyspace{
			{Name: "k1", Tables: []*sdk.Table{{
				Name:       "t1",
				PrimaryKey: &sdk.PrimaryKey{PartitionKey: []string{"c1"}},
				Columns: []*sdk.Column{
					{Name: "c1", DataType: gocqlhelpers.NewTypeText()},
					{Name: "c2", DataType: gocqlhelpers.NewTypeSet(gocqlhelpers.NewTypeTimestamp())},
				},
			}}},
			{Name: defaultKeyspaceName},
		},
	}
	require.True(t, reflect.DeepEqual(expected, s1))
}
