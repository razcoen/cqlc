package config

import (
	"bytes"
	"testing"

	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Run("missing cqls", func(t *testing.T) {
		config := &Config{}
		require.ErrorContains(t, config.Validate(), "validation", "cql")
		config = &Config{CQL: []*CQL{}}
		require.ErrorContains(t, config.Validate(), "validation", "cql")
	})
	t.Run("missing queries", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema: "schema",
				Gen:    &CQLGen{},
			},
		}}
		require.ErrorContains(t, config.Validate(), "validation", "queries")
	})
	t.Run("missing schema", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Queries: "queries",
				Gen:     &CQLGen{},
			},
		}}
		require.ErrorContains(t, config.Validate(), "validation", "schema")
	})
	t.Run("missing gen", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema:  "schema",
				Queries: "queries",
			},
		}}
		require.ErrorContains(t, config.Validate(), "validation", "gen")
	})
	t.Run("missing gen go package", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema:  "schema",
				Queries: "queries",
				Gen: &CQLGen{
					Go: &golang.Options{
						Out: "out",
					},
				},
			},
		}}
		require.ErrorContains(t, config.Validate(), "validation", "package")
	})
	t.Run("missing gen go out", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema:  "schema",
				Queries: "queries",
				Gen: &CQLGen{
					Go: &golang.Options{
						Package: "package",
					},
				},
			},
		}}
		require.ErrorContains(t, config.Validate(), "validation", "out")
	})
	t.Run("valid config", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema:  "schema",
				Queries: "queries",
				Gen:     &CQLGen{},
			},
		}}
		require.NoError(t, config.Validate())
	})
	t.Run("valid config with gen go", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema:  "schema",
				Queries: "queries",
				Gen: &CQLGen{
					Go: &golang.Options{
						Package: "package",
						Out:     "out",
					},
				},
			},
		}}
		require.NoError(t, config.Validate())
	})
}

func TestParseConfig(t *testing.T) {
	t.Run("invalid config", func(t *testing.T) {
		buf := bytes.Buffer{}
		str := `a:
  - b: "b"
    c: "c"
`
		buf.WriteString(str)
		config, err := ParseConfig(&buf)
		// TODO: Validate the error message
		require.Error(t, err)
		require.Nil(t, config)
	})
	t.Run("valid config", func(t *testing.T) {
		buf := bytes.Buffer{}
		str := `cql:
  - schema: "schema"
    queries: "queries"
    gen:
      go:
        package: "package"
        out: "out"
`
		buf.WriteString(str)
		config, err := ParseConfig(&buf)
		require.NoError(t, err)
		require.Equal(t, &Config{
			CQL: []*CQL{{
				Queries: "queries",
				Schema:  "schema",
				Gen:     &CQLGen{Go: &golang.Options{Package: "package", Out: "out"}}},
			},
		}, config)
	})
}
