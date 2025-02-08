package cqlc

import (
	"testing"

	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/razcoen/cqlc/pkg/cqlc/internal/tools"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("partial config", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Gen: &config.CQLGen{},
				},
			},
		})
		require.Error(t, err)
	})
	t.Run("missing schema file", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Schema:  "nonexistingfile",
					Queries: "internal/testdata/basic_queries.cql",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "basic",
							Out:     "internal/testgen/basic",
						},
					},
				},
			},
		})
		require.Error(t, err)
	})
	t.Run("missing queries file", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Schema:  "internal/testdata/basic_schema.cql",
					Queries: "nonexistingfile",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "basic",
							Out:     "internal/testgen/basic",
						},
					},
				},
			},
		})
		require.Error(t, err)
	})
	t.Run("missing gen go", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Queries: "internal/testdata/basic_queries.cql",
					Schema:  "internal/testdata/basic_schema.cql",
					Gen:     &config.CQLGen{},
				},
			},
		})
		require.ErrorContains(t, err, "golang generation config is required: only golang support")
	})
	t.Run("basic", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Queries: "internal/testdata/basic_queries.cql",
					Schema:  "internal/testdata/basic_schema.cql",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "basic",
							Out:     "internal/testgen/basic",
						},
					},
				},
			},
		})
		require.NoError(t, err)
		formatDirAndExpectForNoDiff(t, "internal/testgen/basic")
	})
	t.Run("keyspaced query", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Queries: "internal/testdata/keyspaced_queries.cql",
					Schema:  "internal/testdata/keyspaced_schema.cql",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "keyspaced",
							Out:     "internal/testgen/keyspaced",
						},
					},
				},
			},
		})
		require.NoError(t, err)
		formatDirAndExpectForNoDiff(t, "internal/testgen/keyspaced")
	})
	t.Run("partially keyspaced queries", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Queries: "internal/testdata/partially_keyspaced_queries.cql",
					Schema:  "internal/testdata/keyspaced_schema.cql",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "partiallykeyspaced",
							Out:     "internal/testgen/partiallykeyspaced",
						},
					},
				},
			},
		})
		require.Error(t, err)
		require.NoDirExists(t, "internal/testgen/partiallykeyspaced")
	})
	t.Run("invalid queries", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Queries: "internal/testdata/invalid_queries.cql",
					Schema:  "internal/testdata/basic_schema.cql",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "invalidqueries",
							Out:     "internal/testgen/invalidqueries",
						},
					},
				},
			},
		})
		require.Error(t, err)
		require.NoDirExists(t, "internal/testgen/invalidqueries")
	})
	t.Run("invalid schema", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Queries: "internal/testdata/basic_queries.cql",
					Schema:  "internal/testdata/invalid_schema.cql",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "invalidschema",
							Out:     "internal/testgen/invalidschema",
						},
					},
				},
			},
		})
		require.Error(t, err)
		require.NoDirExists(t, "internal/testgen/invalidschema")
	})
}

func formatDirAndExpectForNoDiff(t *testing.T, dir string) {
	require.NoError(t, tools.Goimports(dir))
	diff, err := tools.GitStatus(dir)
	require.NoError(t, err)
	require.Empty(t, diff, "newly generated code differs from the existing: commit the new generated code first")
}
