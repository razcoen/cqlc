package testcassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

// SessionWrapper wraps a gocql.Session and a keyspace name.
// It is used to manage the lifecycle of a session and a keyspace.
type SessionWrapper struct {
	Session  *gocql.Session
	Keyspace string

	dropKeyspaceOnCleanup bool
}

// Close closes the session and drops the keyspace if it was created by the wrapper.
func (s *SessionWrapper) Close() error {
	defer s.Session.Close()
	if !s.dropKeyspaceOnCleanup {
		return nil
	}
	if err := dropKeyspace(s.Session, s.Keyspace); err != nil {
		return fmt.Errorf("drop keyspace %s: %w", s.Keyspace, err)
	}
	return nil
}

// ConnectWithRandomKeyspace creates a new session with a random keyspace.
func ConnectWithRandomKeyspace() (*SessionWrapper, error) {
	adminSession, err := createSession("system")
	if err != nil {
		return nil, fmt.Errorf("create admin session: %w", err)
	}
	defer adminSession.Close()
	keyspace, err := createRandomKeyspace(adminSession)
	if err != nil {
		return nil, fmt.Errorf("create random keyspace: %w", err)
	}
	session, err := createSession(keyspace)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	return &SessionWrapper{
		Session:               session,
		Keyspace:              keyspace,
		dropKeyspaceOnCleanup: true,
	}, nil
}

// Connect creates a new session with the given keyspace.
func Connect(keyspace string) (*SessionWrapper, error) {
	session, err := createSession(keyspace)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	return &SessionWrapper{
		Session:               session,
		Keyspace:              keyspace,
		dropKeyspaceOnCleanup: false,
	}, nil
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
