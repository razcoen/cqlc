package testcassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func NewRandomKeyspaceSession(t *testing.T) (session *gocql.Session, keyspace string) {
	adminSession := establishAdminSession(t)
	keyspace = createRandomKeyspace(t, adminSession)
	t.Cleanup(func() {
		dropKeyspace(t, adminSession, keyspace)
		adminSession.Close()
	})
	session = establishSession(t, keyspace)
	return session, keyspace
}

func establishAdminSession(t *testing.T) *gocql.Session {
	return establishSession(t, "system")
}

func establishSession(t *testing.T, keyspace string) *gocql.Session {
	sleep := time.Second
	timeout := time.Minute
	deadline := time.Now().Add(timeout)
	var err error
	var session *gocql.Session
	for time.Now().Before(deadline) {
		session, err = createSession(keyspace)
		if err == nil {
			return session
		}
		time.Sleep(sleep)
	}
	require.NoError(t, err)
	return nil
}

func createSession(keyspace string) (*gocql.Session, error) {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = keyspace
	return cluster.CreateSession()
}

func createRandomKeyspace(t *testing.T, session *gocql.Session) string {
	keyspaceID, err := uuid.NewRandom()
	require.NoError(t, err)
	keyspace := fmt.Sprintf("test%x", keyspaceID[:])
	stmt := fmt.Sprintf("CREATE KEYSPACE %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}", keyspace)
	err = session.Query(stmt).Exec()
	require.NoError(t, err)
	return keyspace
}

func dropKeyspace(t *testing.T, session *gocql.Session, keyspace string) {
	stmt := fmt.Sprintf("DROP KEYSPACE %s", keyspace)
	err := session.Query(stmt).Exec()
	require.NoError(t, err)
}
