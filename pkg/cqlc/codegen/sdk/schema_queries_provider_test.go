package sdk

import (
	"testing"

	"github.com/razcoen/cqlc/pkg/cqlc/gocqlhelpers"
	"github.com/stretchr/testify/require"
)

func TestCompileSchemaWithQueries(t *testing.T) {
	schema := &Schema{
		Keyspaces: []*Keyspace{{
			Name: "ks1",
			Tables: []*Table{
				{
					Name: "tbl1",
					Columns: []*Column{
						{Name: "col1", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "col2", DataType: gocqlhelpers.NewTypeTime()},
						{Name: "col3", DataType: gocqlhelpers.NewTypeInt()}},
					PrimaryKey: &PrimaryKey{
						PartitionKey:  []string{"col1"},
						ClusteringKey: []string{"col3"},
					}},
			}}}}
	requireSchemaValidity := func(t *testing.T, p *SchemaQueriesProvider) {
		require.Equal(t, schema, p.Schema())
		require.True(t, p.HasTable("ks1", "tbl1"))
		require.False(t, p.HasTable("ks2", "tbl1"))
		require.False(t, p.HasTable("ks1", "tbl2"))
	}
	newValidInsertQuery := func() *Query {
		return &Query{
			FuncName:    "BasicInsertOne",
			Annotations: []string{"exec"},
			Stmt:        "INSERET INTO ks1.tbl1 (col1,col2,col3) VALUES (?,?,?);",
			Params:      []string{"col1", "col2", "col3"},
			Selects:     []string{},
			Table:       "tbl1",
			Keyspace:    "ks1",
		}
	}
	t.Run("no queries", func(t *testing.T) {
		p, err := CompileSchemaWithQueries(schema, nil)
		require.NoError(t, err)
		requireSchemaValidity(t, p)
		require.Empty(t, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("valid insert query", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{validInsertQuery})
		require.NoError(t, err)
		requireSchemaValidity(t, p)
		require.Len(t, p.ListTableQueries("ks1", "tbl1"), 1)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
		require.Len(t, p.ListTableQueries("ks2", "tbl1"), 0)
		require.Len(t, p.ListTableQueries("ks1", "tbl2"), 0)
	})
	t.Run("query with invalid param", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "InsertOne",
				Annotations: []string{"exec"},
				Stmt:        "INSERET INTO ks1.tbl1 (col1,col2,col4) VALUES (?,?,?);",
				Params:      []string{"col1", "col2", "col4"},
				Selects:     []string{},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
		})
		require.ErrorContains(t, err, `query "InsertOne": parametrized column "col4" does not exist in table "ks1.tbl1"`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("query with invalid keyspace", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "InsertOne",
				Annotations: []string{"exec"},
				Stmt:        "INSERET INTO ks2.tbl1 (col1,col2,col3) VALUES (?,?,?);",
				Params:      []string{"col1", "col2", "col3"},
				Selects:     []string{},
				Table:       "tbl1",
				Keyspace:    "ks2",
			},
		})
		require.ErrorContains(t, err, `query "InsertOne": keyspace "ks2" does not exist in schema`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("query with invalid table", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "InsertOne",
				Annotations: []string{"exec"},
				Stmt:        "INSERET INTO ks1.tbl2 (col1,col2,col3) VALUES (?,?,?);",
				Params:      []string{"col1", "col2", "col3"},
				Selects:     []string{},
				Table:       "tbl2",
				Keyspace:    "ks1",
			},
		})
		require.ErrorContains(t, err, `query "InsertOne": table "tbl2" does not exist in keyspace "ks1"`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("query with invalid select", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "FindOne",
				Annotations: []string{"one"},
				Stmt:        "SELECT col4 FROM ks1.tbl1 WHERE col1 = ?;",
				Params:      []string{"col1"},
				Selects:     []string{"col4"},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
		})
		require.ErrorContains(t, err, `query "FindOne": selected column "col4" does not exist in table "ks1.tbl1"`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("query with invalid astrix select", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "FindOne",
				Annotations: []string{"one"},
				Stmt:        "SELECT *, col1 FROM ks1.tbl1 WHERE col1 = ?;",
				Params:      []string{"col1"},
				Selects:     []string{"*", "col1"},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
		})
		require.ErrorContains(t, err, `query "FindOne": cannot select both * and other columns: choose either * or specific columns`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("query with astrix select", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "FindOne",
				Annotations: []string{"one"},
				Stmt:        "SELECT * FROM ks1.tbl1 WHERE col1 = ?;",
				Params:      []string{"col1"},
				Selects:     []string{"*"},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
		})
		require.NoError(t, err)
		requireSchemaValidity(t, p)
		require.Len(t, p.ListTableQueries("ks1", "tbl1"), 2)
		require.Equal(t, validInsertQuery, p.ListTableQueries("ks1", "tbl1")[0])
		require.Equal(
			t,
			Query{
				FuncName:    "FindOne",
				Annotations: []string{"one"},
				Stmt:        "SELECT * FROM ks1.tbl1 WHERE col1 = ?;",
				Params:      []string{"col1"},
				// Expect that the selections will be modified according to cassandra natural order
				Selects:  []string{"col1", "col3", "col2"},
				Table:    "tbl1",
				Keyspace: "ks1",
			},
			*p.ListTableQueries("ks1", "tbl1")[1],
		)
	})
	t.Run("query with invalid function name", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "?InsertOne",
				Annotations: []string{"exec"},
				Stmt:        "INSERET INTO ks1.tbl1 (col1,col2,col3) VALUES (?,?,?);",
				Params:      []string{"col1", "col2", "col3"},
				Selects:     []string{},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
		})
		require.ErrorContains(t, err, `query "?InsertOne": invalid query function name selected "?InsertOne": string must follow the regexp "^[A-Za-z_][A-Za-z0-9_:<>~]*$"`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("query without annotations", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "InsertOne",
				Annotations: []string{},
				Stmt:        "INSERET INTO ks1.tbl1 (col1,col2,col3) VALUES (?,?,?);",
				Params:      []string{"col1", "col2", "col3"},
				Selects:     []string{},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
		})
		require.ErrorContains(t, err, `query "InsertOne": missing annotations: use of of the following: [exec one many batch]`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
	t.Run("multiple invalid queries", func(t *testing.T) {
		validInsertQuery := newValidInsertQuery()
		p, err := CompileSchemaWithQueries(schema, Queries{
			validInsertQuery,
			{
				FuncName:    "InsertOne",
				Annotations: []string{"exec"},
				Stmt:        "INSERET INTO ks1.tbl1 (col1,col2,col4) VALUES (?,?,?);",
				Params:      []string{"col1", "col2", "col4"},
				Selects:     []string{},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
			{
				FuncName:    "?InsertOne",
				Annotations: []string{"exec"},
				Stmt:        "INSERET INTO ks1.tbl1 (col1,col2,col3) VALUES (?,?,?);",
				Params:      []string{"col1", "col2", "col3"},
				Selects:     []string{},
				Table:       "tbl1",
				Keyspace:    "ks1",
			},
		})
		require.ErrorContains(t, err, `query "InsertOne": parametrized column "col4" does not exist in table "ks1.tbl1"`)
		require.ErrorContains(t, err, `query "?InsertOne": invalid query function name selected "?InsertOne": string must follow the regexp "^[A-Za-z_][A-Za-z0-9_:<>~]*$"`)
		requireSchemaValidity(t, p)
		require.ElementsMatch(t, Queries{validInsertQuery}, p.ListTableQueries("ks1", "tbl1"))
	})
}
