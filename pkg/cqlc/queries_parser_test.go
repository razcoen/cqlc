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
			expectedErr: false,
			expectedQueries: []*Query{
				{
					Stmt:        "SELECT * FROM users",
					Table:       "users",
					Selects:     []string{"*"},
					FuncName:    "ListUsers",
					Annotations: []string{"many"},
				}},
		},
		{
			query: `
-- name: ListUserNamesOfAge :many
SELECT id, name FROM users WHERE age > ?;
`,
			expectedErr: false,
			expectedQueries: []*Query{
				{
					Stmt:        "SELECT id, name FROM users WHERE age > ?",
					Table:       "users",
					Selects:     []string{"id", "name"},
					Params:      []string{"age"},
					FuncName:    "ListUserNamesOfAge",
					Annotations: []string{"many"},
				}},
		},
		{
			query: `
-- name: ListUserNamesOfAgeOrderByName :many
SELECT id, name FROM users WHERE age > ? ORDER BY name ASC;
`,
			expectedQueries: []*Query{
				{
					Stmt:        "SELECT id, name FROM users WHERE age > ? ORDER BY name ASC",
					Table:       "users",
					Selects:     []string{"id", "name"},
					Params:      []string{"age"},
					FuncName:    "ListUserNamesOfAgeOrderByName",
					Annotations: []string{"many"},
				}},
		},
		{
			query: `
-- name: ListActiveUserIDsOfAge :many
SELECT id FROM users WHERE age > ? AND active = ?;
`,
			expectedQueries: []*Query{
				{
					Stmt:        "SELECT id FROM users WHERE age > ? AND active = ?",
					Table:       "users",
					Selects:     []string{"id"},
					Params:      []string{"age", "active"},
					FuncName:    "ListActiveUserIDsOfAge",
					Annotations: []string{"many"},
				}},
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
			query: `
-- name: CreateUser :exec
INSERT INTO users (id, name, age) VALUES (?, ?, 10);
`,
			expectedErr: false,
			expectedQueries: []*Query{
				{
					Stmt:        "INSERT INTO users (id, name, age) VALUES (?, ?, 10)",
					Table:       "users",
					Params:      []string{"id", "name"},
					FuncName:    "CreateUser",
					Annotations: []string{"exec"},
				},
			},
		},
		{
			query: `
-- name: CreateUser :exec
INSERT INTO users (id, name) VALUES (?, ?);
`,
			expectedErr: false,
			expectedQueries: []*Query{
				{
					Stmt:        "INSERT INTO users (id, name) VALUES (?, ?)",
					Table:       "users",
					Params:      []string{"id", "name"},
					FuncName:    "CreateUser",
					Annotations: []string{"exec"},
				},
			},
		},
		{
			query: `
-- name: CreateUser :batch
INSERT INTO users (id, active, name) VALUES (?, true, ?);
`,
			expectedErr: false,
			expectedQueries: []*Query{
				{
					Stmt:        "INSERT INTO users (id, active, name) VALUES (?, true, ?)",
					Table:       "users",
					Params:      []string{"id", "name"},
					FuncName:    "CreateUser",
					Annotations: []string{"batch"},
				},
			},
		},
		// Invalid INSERT Queries
		{
			query:       "INSERT INTO users (id, name age) VALUES (?, ?, ?);",
			expectedErr: true,
		},
		{
			query:       "INSERT INTO users (id name, age) VALUES (?, ?, ?);",
			expectedErr: true,
		},
		{
			query:       "INSERT INTO users (id, name, age) VALUE (?, ?, ?);",
			expectedErr: true,
		},

		// Valid DELETE Queries
		{
			query: `
-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
`,
			expectedErr: false,
			expectedQueries: []*Query{
				{
					Stmt:        "DELETE FROM users WHERE id = ?",
					Params:      []string{"id"},
					Table:       "users",
					FuncName:    "DeleteUser",
					Annotations: []string{"exec"},
				},
			},
		},
		{
			query: `
-- name: DeleteUsersOver20ByName :exec
DELETE FROM users WHERE name = ? AND age > 20;
`,
			expectedErr: false,
			expectedQueries: []*Query{
				{
					Stmt:        "DELETE FROM users WHERE name = ? AND age > 20",
					Params:      []string{"name"},
					Table:       "users",
					FuncName:    "DeleteUsersOver20ByName",
					Annotations: []string{"exec"},
				},
			},
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
