package cmd

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"testing"
	"time"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
)

func TestNewVersionCommand(t *testing.T) {
	noopLogger := slog.New(slog.NewTextHandler(&bytes.Buffer{}, &slog.HandlerOptions{}))
	t.Run("flags", func(t *testing.T) {
		cmd := NewVersionCommand(noopLogger, &buildinfo.BuildInfo{
			Version:   "v1.0.0",
			Commit:    "7d23a9c24c9e683f76ddb01e38a0404838628cb0",
			Time:      time.Now(),
			GoVersion: "go1.16.3",
		})
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			require.Contains(t, []string{"format"}, f.Name)
		})
	})
	t.Run("format text", func(t *testing.T) {
		cmd := NewVersionCommand(noopLogger, &buildinfo.BuildInfo{
			Version:   "v1.0.0",
			Commit:    "7d23a9c24c9e683f76ddb01e38a0404838628cb0",
			Time:      time.Now(),
			GoVersion: "go1.16.3",
		})
		buf := &bytes.Buffer{}
		cmd.SetOut(buf)
		err := cmd.Execute()
		require.NoError(t, err)
		require.Contains(t, buf.String(), "cqlc version v1.0.0")
		require.Contains(t, buf.String(), "7d23a9c24c9e683f76ddb01e38a0404838628cb0")
		require.Contains(t, buf.String(), "go1.16.3")
	})
	t.Run("format json", func(t *testing.T) {
		bi := &buildinfo.BuildInfo{
			Version:   "v1.0.0",
			Commit:    "7d23a9c24c9e683f76ddb01e38a0404838628cb0",
			Time:      time.Now(),
			GoVersion: "go1.16.3",
		}
		cmd := NewVersionCommand(noopLogger, bi)
		require.NoError(t, cmd.Flags().Set("format", "json"))
		buf := &bytes.Buffer{}
		cmd.SetOut(buf)
		err := cmd.Execute()
		require.NoError(t, err)
		var out buildinfo.BuildInfo
		err = json.Unmarshal(buf.Bytes(), &out)
		require.NoError(t, err)
		// Time cannot be compared indirectly.
		require.Equal(t, bi.Time.Unix(), out.Time.Unix())
		out.Time = bi.Time // Reset Time to avoid comparison error.
		require.Equal(t, bi, &out)
	})
}
