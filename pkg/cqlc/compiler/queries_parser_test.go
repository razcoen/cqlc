package compiler

import (
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueriesParser(t *testing.T) {
	// List of test cases for schema-related queries
	tests := []struct {
		name            string
		query           string
		expectedErr     bool
		expectedQueries []*sdk.Query
	}{
		// Valid SELECT Queries
		{
			name: "select all",
			query: `
-- name: ListUsers :many
SELECT * FROM users;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "SELECT * FROM users;",
					Table:       "users",
					Selects:     []string{"*"},
					FuncName:    "ListUsers",
					Annotations: []string{"many"},
				}},
		},
		{
			name: "select all and specify keyspace",
			query: `
-- name: ListUsers :many
SELECT * FROM auth.users;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "SELECT * FROM auth.users;",
					Table:       "users",
					Keyspace:    "auth",
					Selects:     []string{"*"},
					FuncName:    "ListUsers",
					Annotations: []string{"many"},
				}},
		},
		{
			name: "select columns with parameters",
			query: `
-- name: ListUserNamesOfAge :many
SELECT id, name FROM users WHERE age > ?;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "SELECT id, name FROM users WHERE age > ?;",
					Table:       "users",
					Selects:     []string{"id", "name"},
					Params:      []string{"age"},
					FuncName:    "ListUserNamesOfAge",
					Annotations: []string{"many"},
				}},
		},
		{
			name: "select columns with parameters use ordering",
			query: `
-- name: ListUserNamesOfAgeOrderByName :many
SELECT id, name FROM users WHERE age > ? ORDER BY name ASC;
`,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "SELECT id, name FROM users WHERE age > ? ORDER BY name ASC;",
					Table:       "users",
					Selects:     []string{"id", "name"},
					Params:      []string{"age"},
					FuncName:    "ListUserNamesOfAgeOrderByName",
					Annotations: []string{"many"},
				}},
		},
		{
			name: "select columns with multiple parameters",
			query: `
-- name: ListActiveUserIDsOfAge :many
SELECT id FROM users WHERE age > ? AND active = ?;
`,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "SELECT id FROM users WHERE age > ? AND active = ?;",
					Table:       "users",
					Selects:     []string{"id"},
					Params:      []string{"age", "active"},
					FuncName:    "ListActiveUserIDsOfAge",
					Annotations: []string{"many"},
				}},
		},
		// Invalid SELECT Queries
		{
			name:        "select without columns",
			query:       "SELECT FROM users;",
			expectedErr: true,
		},
		{
			name:        "select columns without comma",
			query:       "SELECT id name FROM users;",
			expectedErr: true,
		},
		{
			name:        "select misplace table name",
			query:       "SELECT * users WHERE age > 25;",
			expectedErr: true,
		},

		// Valid INSERT Queries
		{
			name: "insert with parameters and int constant",
			query: `
-- name: CreateUser :exec
INSERT INTO users (id, name, age) VALUES (?, ?, 10);
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "INSERT INTO users (id, name, age) VALUES (?, ?, 10);",
					Table:       "users",
					Params:      []string{"id", "name"},
					FuncName:    "CreateUser",
					Annotations: []string{"exec"},
				},
			},
		},
		{
			name: "insert with parameters",
			query: `
-- name: CreateUser :exec
INSERT INTO users (id, name) VALUES (?, ?);
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "INSERT INTO users (id, name) VALUES (?, ?);",
					Table:       "users",
					Params:      []string{"id", "name"},
					FuncName:    "CreateUser",
					Annotations: []string{"exec"},
				},
			},
		},
		{
			name: "insert with parameters and bool constant",
			query: `
-- name: CreateUser :batch
INSERT INTO users (id, active, name) VALUES (?, true, ?);
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "INSERT INTO users (id, active, name) VALUES (?, true, ?);",
					Table:       "users",
					Params:      []string{"id", "name"},
					FuncName:    "CreateUser",
					Annotations: []string{"batch"},
				},
			},
		},
		// Invalid INSERT Queries
		{
			name:        "insert values missing comma 1",
			query:       "INSERT INTO users (id, name age) VALUES (?, ?, ?);",
			expectedErr: true,
		},
		{
			name:        "insert values missing comma 2",
			query:       "INSERT INTO users (id name, age) VALUES (?, ?, ?);",
			expectedErr: true,
		},
		{
			name:        "insert values typo as value",
			query:       "INSERT INTO users (id, name, age) VALUE (?, ?, ?);",
			expectedErr: true,
		},

		// Valid DELETE Queries
		{
			name: "delete with parameter",
			query: `
-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "DELETE FROM users WHERE id = ?;",
					Params:      []string{"id"},
					Table:       "users",
					FuncName:    "DeleteUser",
					Annotations: []string{"exec"},
				},
			},
		},
		{
			name: "delete with parameter and constant",
			query: `
-- name: DeleteUsersOver20ByName :exec
DELETE FROM users WHERE name = ? AND age > 20;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        "DELETE FROM users WHERE name = ? AND age > 20;",
					Params:      []string{"name"},
					Table:       "users",
					FuncName:    "DeleteUsersOver20ByName",
					Annotations: []string{"exec"},
				},
			},
		},
		// Invalid DELETE Queries
		{
			name:        "delete without from keyword",
			query:       "DELETE users WHERE id = 2;",
			expectedErr: true,
		},
		{
			name:        "delete without table",
			query:       "DELETE FROM WHERE id = 3;",
			expectedErr: true,
		},
		{
			name:        "delete without table or where",
			query:       "DELETE FROM users id = 4;",
			expectedErr: true,
		},

		// Edge cases
		{
			name: "ignore comments other than last",
			query: `
-- This is a possible comment.
/*
 A comment can also include keywords like TIMESTAMP and ;
 */
// This is also be a valid comment.
-- name: FindUsers :many
SELECT * FROM users WHERE email = ? ALLOW FILTERING;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        `SELECT * FROM users WHERE email = ? ALLOW FILTERING;`,
					Params:      []string{"email"},
					Selects:     []string{"*"},
					Table:       "users",
					FuncName:    "FindUsers",
					Annotations: []string{"many"},
				},
			},
		},
		{
			name: "ignore empty lines after comment",
			query: `
-- name: FindUsers :many


SELECT * FROM users WHERE email = ? ALLOW FILTERING;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        `SELECT * FROM users WHERE email = ? ALLOW FILTERING;`,
					Params:      []string{"email"},
					Selects:     []string{"*"},
					Table:       "users",
					FuncName:    "FindUsers",
					Annotations: []string{"many"},
				},
			},
		},
		{
			name: "multi line astQueryAnalysis",
			query: `
-- name: FindUsers :many
SELECT * FROM users
WHERE email = ?
ALLOW FILTERING;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        `SELECT * FROM users WHERE email = ? ALLOW FILTERING;`,
					Params:      []string{"email"},
					Selects:     []string{"*"},
					Table:       "users",
					FuncName:    "FindUsers",
					Annotations: []string{"many"},
				},
			},
		},
		{
			name: "remove extra spaces",
			query: `
-- name: FindUsers :many
SELECT * FROM users  WHERE email =     ? ALLOW FILTERING;
`,
			expectedErr: false,
			expectedQueries: []*sdk.Query{
				{
					Stmt:        `SELECT * FROM users WHERE email = ? ALLOW FILTERING;`,
					Params:      []string{"email"},
					Selects:     []string{"*"},
					Table:       "users",
					FuncName:    "FindUsers",
					Annotations: []string{"many"},
				},
			},
		},
	}
	parser := NewQueriesParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queries, err := parser.Parse(tt.query)
			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error: %v, but got: %v for astQueryAnalysis: %s", tt.expectedErr, err, tt.query)
			}
			if err != nil || tt.expectedErr {
				return
			}
			require.ElementsMatch(t, tt.expectedQueries, queries)
		})
	}
}
