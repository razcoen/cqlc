package gocqlc

import (
	"testing"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/require"
)

func TestBatchOption(t *testing.T) {
	t.Run("with batch type", func(t *testing.T) {
		b := gocql.Batch{}
		require.Equal(t, gocql.LoggedBatch, b.Type)
		opt := WithBatchType(gocql.CounterBatch)
		opt.Apply(&b)
		require.Equal(t, gocql.CounterBatch, b.Type)
	})
}
