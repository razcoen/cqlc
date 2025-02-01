package golang

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOptionsValidate(t *testing.T) {
	t.Run("missing package", func(t *testing.T) {
		opts := Options{Out: "out"}
		require.ErrorContains(t, opts.Validate(), "validation", "package")
	})
	t.Run("missing out", func(t *testing.T) {
		opts := Options{Package: "package"}
		require.ErrorContains(t, opts.Validate(), "validation", "out")
	})
	t.Run("valid options", func(t *testing.T) {
		opts := Options{Package: "package", Out: "out"}
		require.NoError(t, opts.Validate())
	})
}
