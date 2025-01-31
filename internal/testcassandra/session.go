package testcassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type SessionWrapper struct {
	Session  *gocql.Session
	Keyspace string

	// adminSession is used to create and drop keyspaces.
	// It should not be used for any other purpose.
	adminSession *gocql.Session
}

func (s *SessionWrapper) Close() error {
	var closeErr error
	if err := dropKeyspace(s.adminSession, s.Keyspace); err != nil {
		closeErr = fmt.Errorf("drop keyspace %s: %w", s.Keyspace, err)
	}
	s.adminSession.Close()
	s.Session.Close()
	return closeErr
}

func ConnectWithRandomKeyspace() (*SessionWrapper, error) {
	adminSession, err := createSession("system")
	if err != nil {
		return nil, fmt.Errorf("create admin session: %w", err)
	}
	keyspace, err := createRandomKeyspace(adminSession)
	if err != nil {
		return nil, fmt.Errorf("create random keyspace: %w", err)
	}
	session, err := createSession(keyspace)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	return &SessionWrapper{Session: session, Keyspace: keyspace, adminSession: adminSession}, nil
}

func createSession(keyspace string) (*gocql.Session, error) {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = keyspace
	return cluster.CreateSession()
}

func createRandomKeyspace(session *gocql.Session) (string, error) {
	keyspaceID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("generate keyspace id: %w", err)
	}
	keyspace := fmt.Sprintf("test%x", keyspaceID[:])
	stmt := fmt.Sprintf("CREATE KEYSPACE %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}", keyspace)
	if err := session.Query(stmt).Exec(); err != nil {
		return "", fmt.Errorf("create keyspace: %w", err)
	}
	return keyspace, nil
}

func dropKeyspace(session *gocql.Session, keyspace string) error {
	stmt := fmt.Sprintf("DROP KEYSPACE %s", keyspace)
	return session.Query(stmt).Exec()
}
