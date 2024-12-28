package cql

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestSchemaBuilder(t *testing.T) {
	t1, err := NewTableBuilder("t1").
		WithColumn("c1", NativeTypeText.IntoDataType()).
		WithColumn("c2", CollectionTypeSet{T: NativeTypeTimestamp.IntoCollectableType()}.IntoDataType()).
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
					{Name: "c1", DataType: NativeTypeText.IntoDataType()},
					{Name: "c2", DataType: CollectionTypeSet{T: NativeTypeTimestamp.IntoCollectableType()}.IntoDataType()},
				},
			}}},
			{Name: defaultKeyspaceName},
		},
	}
	diff := cmp.Diff(expected, s1)
	if diff != "" {
		t.Errorf("returned schema different than expected: %s", diff)
	}
}
