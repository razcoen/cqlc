package buildinfo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoreAndLoad(t *testing.T) {
	t.Run("empty version", func(t *testing.T) {
		t.Cleanup(func() { buildInfo.Store(nil) })
		err := Store(&Flags{})
		require.Error(t, err)
		require.Nil(t, Load())
	})
	t.Run("invalid version", func(t *testing.T) {
		t.Cleanup(func() { buildInfo.Store(nil) })
		err := Store(&Flags{Version: "hello"})
		require.Error(t, err)
		require.Nil(t, Load())
	})
	t.Run("valid version", func(t *testing.T) {
		t.Cleanup(func() { buildInfo.Store(nil) })
		err := Store(&Flags{Version: "v1.0.0"})
		require.NoError(t, err)
		require.Equal(t, "v1.0.0", Load().Version)
	})
	t.Run("snapshot version", func(t *testing.T) {
		t.Cleanup(func() { buildInfo.Store(nil) })
		err := Store(&Flags{Version: "v1.0.0-SNAPSHOT-2021-01-01"})
		require.NoError(t, err)
		require.Equal(t, "v1.0.0-SNAPSHOT-2021-01-01", Load().Version)
	})
}
