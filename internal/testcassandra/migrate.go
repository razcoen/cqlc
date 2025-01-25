package testcassandra

import (
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Exec(t *testing.T, session *gocql.Session, path string) {
	b, err := os.ReadFile(path)
	require.NoError(t, err)
	stmt := string(b)
	err = session.Query(stmt).Exec()
	require.NoError(t, err)
}
