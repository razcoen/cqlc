package testcassandra

import (
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Migrate(t *testing.T, session *gocql.Session, schemaFile string) {
	b, err := os.ReadFile(schemaFile)
	require.NoError(t, err)
	stmt := string(b)
	err = session.Query(stmt).Exec()
	require.NoError(t, err)
}
