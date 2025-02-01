package cqlc

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	noopLogger := slog.New(slog.NewTextHandler(&bytes.Buffer{}, &slog.HandlerOptions{}))
	config := &Config{DisableOutput: true}
	testOptions := []Option{WithLogger(noopLogger), WithConfig(config)}
	t.Run("empty version", func(t *testing.T) {
		flags := &buildinfo.Flags{Version: ""}
		err := Run(append(testOptions, WithBuildFlags(flags))...)
		require.Error(t, err)
	})
	t.Run("valid version", func(t *testing.T) {
		flags := &buildinfo.Flags{Version: "v1.0.0"}
		err := Run(append(testOptions, WithBuildFlags(flags))...)
		require.NoError(t, err)
	})
}
