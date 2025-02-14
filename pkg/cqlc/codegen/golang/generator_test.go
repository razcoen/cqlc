package golang

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	cblog "github.com/charmbracelet/log"
	"github.com/gocql/gocql"
	"github.com/razcoen/cqlc/pkg/cqlc/codegen/sdk"
	"github.com/razcoen/cqlc/pkg/cqlc/gocqlhelpers"
	"github.com/razcoen/cqlc/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGenerator(t *testing.T) {
	// The main purpose for this test is to make sure that the templates creation is compiling succesfully.
	logger := log.NopLogger()
	g := NewGenerator(logger)
	assert.NotNil(t, g.clientTemplate)
	assert.NotNil(t, g.queriesGoTemplate)
	assert.NotNil(t, g.oneQueryTemplate)
	assert.NotNil(t, g.execQueryTemplate)
	assert.Equal(t, logger, g.logger)
}

func TestGenerate(t *testing.T) {
	// TODO: Make this test actually worth something
	g := NewGenerator(log.NopLogger())
	tmpdir, err := os.MkdirTemp("", "cqlc-golang-generator-test")
	require.NoError(t, err)
	t.Cleanup(func() { _ = os.RemoveAll(tmpdir) })
	schema := &sdk.Schema{
		Keyspaces: []*sdk.Keyspace{{
			Name: "ks1",
			Tables: []*sdk.Table{
				{
					Name: "tbl1",
					Columns: []*sdk.Column{
						{Name: "col1", DataType: gocqlhelpers.NewTypeUUID()},
						{Name: "col2", DataType: gocqlhelpers.NewTypeTime()},
						{Name: "col3", DataType: gocqlhelpers.NewTypeInt()}},
					PrimaryKey: &sdk.PrimaryKey{
						PartitionKey:  []string{"col1"},
						ClusteringKey: []string{"col3"},
					}},
			}}}}
	queries := sdk.Queries{{
		FuncName:    "BasicInsertOne",
		Annotations: []string{"exec"},
		Stmt:        "INSERET INTO ks1.tbl1 (col1,col2,col3) VALUES (?,?,?);",
		Params:      []string{"col1", "col2", "col3"},
		Selects:     []string{},
		Table:       "tbl1",
		Keyspace:    "ks1",
	}}
	provider, err := sdk.CompileSchemaWithQueries(schema, queries)
	require.NoError(t, err)
	err = g.Generate(
		&sdk.Context{
			Provider: provider,
			Metadata: &sdk.Metadata{SchemaPath: "schema.cql", QueriesPath: "queries.cql", ConfigPath: "cqlc.yaml", Version: "v1.0.0"},
		},
		&Options{
			Package:  "example",
			Out:      filepath.Join(tmpdir, "example"),
			Defaults: DefaultsOptions{},
		},
	)
	require.NoError(t, err)
}

func TestParseBatchType(t *testing.T) {
	t.Run("fallback", func(t *testing.T) {
		assert.Equal(t, "Logged", parseBatchType("hello", gocql.LoggedBatch, nil))
		assert.Equal(t, "Unlogged", parseBatchType("hello", gocql.UnloggedBatch, nil))
		assert.Equal(t, "Counter", parseBatchType("hello", gocql.CounterBatch, nil))
	})
	t.Run("input with logged fallback", func(t *testing.T) {
		assert.Equal(t, "Unlogged", parseBatchType("unlogged", gocql.LoggedBatch, nil))
		assert.Equal(t, "Counter", parseBatchType("counter", gocql.LoggedBatch, nil))
		assert.Equal(t, "Logged", parseBatchType("logged", gocql.LoggedBatch, nil))
	})
	t.Run("input with unlogged fallback", func(t *testing.T) {
		assert.Equal(t, "Unlogged", parseBatchType("unlogged", gocql.UnloggedBatch, nil))
		assert.Equal(t, "Counter", parseBatchType("counter", gocql.UnloggedBatch, nil))
		assert.Equal(t, "Logged", parseBatchType("logged", gocql.UnloggedBatch, nil))
	})
	t.Run("input with counter fallback", func(t *testing.T) {
		assert.Equal(t, "Unlogged", parseBatchType("unlogged", gocql.CounterBatch, nil))
		assert.Equal(t, "Counter", parseBatchType("counter", gocql.CounterBatch, nil))
		assert.Equal(t, "Logged", parseBatchType("logged", gocql.CounterBatch, nil))
	})
	t.Run("logger", func(t *testing.T) {
		var buf bytes.Buffer
		logger := cblog.New(&buf)
		assert.Equal(t, "Logged", parseBatchType("hello", gocql.LoggedBatch, log.NewCharmbraceletAdapter(logger)))
		assert.Contains(t, buf.String(), `using default batch type "logged": invalid batch type "hello" was provided`)
	})
}
