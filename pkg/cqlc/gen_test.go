package cqlc

import (
	"testing"

	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/razcoen/cqlc/pkg/cqlc/internal/tools"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		err := Generate(&config.Config{
			CQL: []*config.CQL{
				{
					Queries: "internal/testdata/basic_queries.cql",
					Schema:  "internal/testdata/basic_schema.cql",
					Gen: &config.CQLGen{
						Go: &golang.Options{
							Package: "basic",
							Out:     "internal/testdata/gen/basic",
						},
					},
				},
			},
		})
		require.NoError(t, err)
		formatDirAndExpectForNoDiff(t, "internal/testdata/gen/basic")
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
							Out:     "internal/testdata/gen/keyspaced",
						},
					},
				},
			},
		})
		require.NoError(t, err)
		formatDirAndExpectForNoDiff(t, "internal/testdata/gen/keyspaced")
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
							Out:     "internal/testdata/gen/partiallykeyspaced",
						},
					},
				},
			},
		})
		require.Error(t, err)
		require.NoDirExists(t, "internal/testdata/gen/partiallykeyspaced")
	})
}

func formatDirAndExpectForNoDiff(t *testing.T, dir string) {
	require.NoError(t, tools.Goimports(dir))
	diff, err := tools.GitStatus(dir)
	require.NoError(t, err)
	require.Empty(t, diff, "newly generated code differs from the existing: commit the new generated code first")
}
