package config

import (
	"testing"

	"github.com/razcoen/cqlc/pkg/cqlc/codegen/golang"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	t.Run("missing queries", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema: "schema",
				Gen:    &CQLGen{},
			},
		}}
		require.NoError(t, config.Validate())
	})
	t.Run("missing schema", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Queries: "queries",
				Gen:     &CQLGen{},
			},
		}}
		require.NoError(t, config.Validate())
	})
	t.Run("missing gen", func(t *testing.T) {
		config := &Config{CQL: []*CQL{
			{
				Schema:  "schema",
				Queries: "queries",
			},
		}}
		require.NoError(t, config.Validate())
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
		require.NoError(t, config.Validate())
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
		require.NoError(t, config.Validate())
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
