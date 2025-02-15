package gocqlc

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/require"
)

func TestQueryOption(t *testing.T) {
	// Unfortunatly, the other options cannot be tested due to them being private.
	t.Run("with consistency", func(t *testing.T) {
		q := gocql.Query{}
		require.Equal(t, gocql.Any, q.GetConsistency())
		opt := WithConsistency(gocql.Quorum)
		opt.apply(&q)
		require.Equal(t, gocql.Quorum, q.GetConsistency())
	})
}

func TestBatchOption(t *testing.T) {
	t.Run("with batch type", func(t *testing.T) {
		b := gocql.Batch{}
		require.Equal(t, gocql.LoggedBatch, b.Type)
		opt := WithBatchType(gocql.CounterBatch)
		opt.apply(&b)
		require.Equal(t, gocql.CounterBatch, b.Type)
	})
}
