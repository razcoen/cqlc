package cqlc

import (
	"log/slog"
	"testing"

	"github.com/razcoen/cqlc/pkg/log"
	"github.com/stretchr/testify/require"
)

func TestNewGeneratorWithOptions(t *testing.T) {
	t.Run("with logger", func(t *testing.T) {
		logger := log.NewSlogAdapter(slog.Default())
		g := newGenerator(WithLogger(logger))
		require.Equal(t, logger, g.logger)
	})
	t.Run("with config path", func(t *testing.T) {
		g := newGenerator(WithConfigPath("config.path.yaml"))
		require.Equal(t, "config.path.yaml", g.configPath)
	})
}
