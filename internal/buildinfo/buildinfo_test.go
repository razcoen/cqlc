package buildinfo

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBuildInfo(t *testing.T) {
	t.Run("empty version", func(t *testing.T) {
		bi, err := ParseBuildInfo(&Flags{})
		require.Error(t, err)
		require.Nil(t, bi)
	})
	t.Run("invalid version", func(t *testing.T) {
		bi, err := ParseBuildInfo(&Flags{})
		require.Error(t, err)
		require.Nil(t, bi)
	})
	t.Run("valid version", func(t *testing.T) {
		bi, err := ParseBuildInfo(&Flags{Version: "v1.0.0"})
		require.NoError(t, err)
		require.Equal(t, "v1.0.0", bi.Version)
	})
	t.Run("snapshot version", func(t *testing.T) {
		bi, err := ParseBuildInfo(&Flags{Version: "v1.0.0-SNAPSHOT-2021-01-01"})
		require.NoError(t, err)
		require.Equal(t, "v1.0.0-SNAPSHOT-2021-01-01", bi.Version)
	})
}

func TestReadModuleVersion(t *testing.T) {
	t.Run("in test runtime", func(t *testing.T) {
		version, err := ReadModuleVersion()
		require.ErrorContains(t, err, "internal error", "no such dependency or module")
		require.Empty(t, version)
	})
	t.Run("in development runtime", func(t *testing.T) {
		out, err := exec.Command("go", "run", "-tags=test", "./internal/testdata/gorun/read_module_version_main.go").Output()
		require.NoError(t, err)
		require.Equal(t, "(devel)", string(out))
	})
}
