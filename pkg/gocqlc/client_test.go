package gocqlc

import (
	"log/slog"
	"testing"

	"github.com/razcoen/cqlc/internal/testcassandra"
	"github.com/razcoen/cqlc/pkg/log"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Run("nil session", func(t *testing.T) {
		client, err := NewClient(nil)
		require.ErrorIs(t, err, ErrNilSession)
		require.Nil(t, client)
	})
	t.Run("closed session", func(t *testing.T) {
		sw, err := testcassandra.ConnectWithRandomKeyspace()
		require.NoError(t, err, "connect to cassandra into random keyspace")
		require.NoError(t, sw.Close(), "close session")
		require.True(t, sw.Session.Closed(), "session should be closed")
		client, err := NewClient(sw.Session)
		require.ErrorIs(t, err, ErrClosedSession)
		require.Nil(t, client)
	})
	sw, err := testcassandra.ConnectWithRandomKeyspace()
	require.NoError(t, err, "connect to cassandra into random keyspace")
	t.Cleanup(func() {
		require.NoError(t, sw.Close(), "close session")
	})
	t.Run("valid session", func(t *testing.T) {
		client, err := NewClient(sw.Session)
		require.NoError(t, err)
		require.Equal(t, sw.Session, client.Session())
	})
	t.Run("with logger", func(t *testing.T) {
		logger := log.NewSlogAdapter(slog.Default())
		client, err := NewClient(sw.Session, WithLogger(logger))
		require.NoError(t, err)
		require.Equal(t, client.Logger(), logger)
	})
}
