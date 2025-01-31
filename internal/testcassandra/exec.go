package testcassandra

import (
	"fmt"
	"os"

	"github.com/gocql/gocql"
)

func ExecFile(session *gocql.Session, path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}
	stmt := string(b)
	if err := session.Query(stmt).Exec(); err != nil {
		return fmt.Errorf("exec file: %w", err)
	}
	return nil
}
