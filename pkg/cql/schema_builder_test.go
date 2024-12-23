package cql

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSchemaBuilder(t *testing.T) {
	t1, err := NewTableBuilder("t1").
		WithColumn("c1", NewNativeTypeText()).
		WithColumn("c2", NewCollectionTypeSet(NativeTypeTimestamp)).
		Build()
	require.NoError(t, err)
	k1, err := NewKeyspaceBuilder("k1").
		WithTable(t1).
		Build()
	require.NoError(t, err)
	s1, err := NewSchemaBuilder().WithKeyspace(k1).Build()
	require.NoError(t, err)
	expected := &Schema{
		Keyspaces: []*Keyspace{
			{Name: "k1", Tables: []*Table{{
				Name: "t1", Columns: []*Column{
					{Name: "c1", DataType: NewNativeTypeText()},
					{Name: "c2", DataType: NewCollectionTypeSet(NativeTypeTimestamp)},
				},
			}}},
			{Name: defaultKeyspaceName},
		},
	}
	require.True(t, reflect.DeepEqual(expected, s1))
}
