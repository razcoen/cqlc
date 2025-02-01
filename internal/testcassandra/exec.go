package testcassandra

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocql/gocql"
)

// ExecFile reads the file at path and executes each statement in the file.
func ExecFile(session *gocql.Session, path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}
	stmts := string(b)
	// TODO: Support comments that include semicolons.
	for _, stmt := range strings.Split(stmts, ";") {
		stmt := strings.TrimSpace(stmt)
		if len(stmt) == 0 {
			continue
		}
		if err := session.Query(stmt).Exec(); err != nil {
			return fmt.Errorf("exec stmt: %w", err)
		}
	}
	return nil
}
