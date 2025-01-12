package cqlc

import (
	"bytes"
	"github.com/razcoen/cqlc/pkg/gocqlhelpers"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoGenerator(t *testing.T) {
	tests := []struct {
		name           string
		keyspace       *Keyspace
		packageName    string
		expectedErr    bool
		expectedOutput string
	}{
		{
			name:        "Empty Keyspace",
			keyspace:    &Keyspace{Name: "empty_keyspace"},
			packageName: "empty",
			expectedErr: false,
			expectedOutput: `// Code generated by cqlc. DO NOT EDIT.

package empty
`,
		},
		{
			name: "Keyspace with Single Table",
			keyspace: &Keyspace{
				Name: "single_table_keyspace",
				Tables: []*Table{
					{
						Name: "users",
						Columns: []*Column{
							{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
							{Name: "name", DataType: gocqlhelpers.NewTypeText()},
							{Name: "age", DataType: gocqlhelpers.NewTypeInt()},
						},
					},
				},
			},
			packageName: "single",
			expectedErr: false,
			expectedOutput: `// Code generated by cqlc. DO NOT EDIT.

package single

import (
	"github.com/gocql/gocql"
)

// Table: users
type User struct {
	ID   gocql.UUID
	Name string
	Age  int
}
`,
		},
		{
			name: "Keyspace with Multiple Tables",
			keyspace: &Keyspace{
				Name: "multi_table_keyspace",
				Tables: []*Table{
					{
						Name: "users",
						Columns: []*Column{
							{Name: "id", DataType: gocqlhelpers.NewTypeUUID()},
							{Name: "name", DataType: gocqlhelpers.NewTypeText()},
						},
					},
					{
						Name: "orders",
						Columns: []*Column{
							{Name: "order_id", DataType: gocqlhelpers.NewTypeUUID()},
							{Name: "user_id", DataType: gocqlhelpers.NewTypeUUID()},
							{Name: "amount", DataType: gocqlhelpers.NewTypeDecimal()},
							{Name: "created_at", DataType: gocqlhelpers.NewTypeTimestamp()},
						},
					},
				},
			},
			packageName: "multi",
			expectedErr: false,
			expectedOutput: `// Code generated by cqlc. DO NOT EDIT.

package multi

import (
	"github.com/gocql/gocql"
	"gopkg.in/inf.v0"
	"time"
)

// Table: users
type User struct {
	ID   gocql.UUID
	Name string
}

// Table: orders
type Order struct {
	OrderID   gocql.UUID
	UserID    gocql.UUID
	Amount    *inf.Dec
	CreatedAt time.Time
}
`,
		},
	}

	// TODO: Test all data types
	gg, err := newGoGenerator()
	require.NoError(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			// TODO: Assert response
			_, err := gg.generateKeyspaceStructs(&generateKeyspaceStructsRequest{
				keyspace:    tt.keyspace,
				packageName: tt.packageName,
				out:         buf,
			})

			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedOutput, buf.String())
			}
		})
	}
}
