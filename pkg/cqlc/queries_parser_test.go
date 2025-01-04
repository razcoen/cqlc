package cqlc

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueriesParser(t *testing.T) {
	// List of test cases for schema-related queries
	tests := []struct {
		query           string
		expectedErr     bool
		expectedQueries []*Query
	}{
		// Valid SELECT Queries
		{
			query: `
-- name: ListUsers :many
SELECT * FROM users;
`,
			expectedErr:     false,
			expectedQueries: []*Query{{Stmt: "SELECT * FROM users", Selects: []string{"*"}, FuncName: "ListUsers", Annotations: []string{"many"}}},
		},
		{
			query:           "SELECT id, name FROM users WHERE age > ?;",
			expectedErr:     false,
			expectedQueries: []*Query{{Stmt: "SELECT id, name FROM users WHERE age > ?", Selects: []string{"id", "name"}, Params: []string{"age"}}},
		},
		{
			query:           "SELECT id, name FROM users WHERE age > ? ORDER BY name ASC;",
			expectedQueries: []*Query{{Stmt: "SELECT id, name FROM users WHERE age > ? ORDER BY name ASC", Selects: []string{"id", "name"}, Params: []string{"age"}}},
		},
		{
			query:           "SELECT id FROM users WHERE age > ? AND active = ?;",
			expectedQueries: []*Query{{Stmt: "SELECT id FROM users WHERE age > ? AND active = ?", Selects: []string{"id"}, Params: []string{"age", "active"}}},
		},
		// Invalid SELECT Queries
		{
			query:       "SELECT FROM users;",
			expectedErr: true,
		},
		{
			query:       "SELECT id name FROM users;",
			expectedErr: true,
		},
		{
			query:       "SELECT * users WHERE age > 25;",
			expectedErr: true,
		},

		// Valid INSERT Queries
		{
			query:       "INSERT INTO users (id, name, age) VALUES (1, 'Alice', 30);",
			expectedErr: false,
		},
		{
			query:       "INSERT INTO users (id, name) VALUES (2, 'Bob');",
			expectedErr: false,
		},
		{
			query:       "INSERT INTO users (id, name, active) VALUES (3, 'Charlie', true);",
			expectedErr: false,
		},
		// Invalid INSERT Queries
		{
			query:       "INSERT INTO users (id, name age) VALUES (4, 'Dave', 40);",
			expectedErr: true,
		},
		{
			query:       "INSERT INTO users (id name, age) VALUES (5, 'Eve', 25);",
			expectedErr: true,
		},
		{
			query:       "INSERT INTO users (id, name, age) VALUE (6, 'Frank', 35);",
			expectedErr: true,
		},

		// Valid DELETE Queries
		{
			query:       "DELETE FROM users WHERE id = 1;",
			expectedErr: false,
		},
		{
			query:       "DELETE FROM users WHERE name = 'Alice';",
			expectedErr: false,
		},
		{
			query:       "DELETE FROM users WHERE age > 30;",
			expectedErr: false,
		},
		// Invalid DELETE Queries
		{
			query:       "DELETE users WHERE id = 2;",
			expectedErr: true,
		},
		{
			query:       "DELETE FROM WHERE id = 3;",
			expectedErr: true,
		},
		{
			query:       "DELETE FROM users id = 4;",
			expectedErr: true,
		},
	}
	parser := NewQueriesParser()
	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			queries, err := parser.Parse(tt.query)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, but got: %v for query: %s", tt.expectedErr, err, tt.query)
			}
			if err != nil || tt.expectedErr {
				return
			}
			require.ElementsMatch(t, tt.expectedQueries, queries)
		})
	}
}
