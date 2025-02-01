package buildinfo

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/mod/semver"
)

// BuildInfo contains information about the build
type BuildInfo struct {
	// Version is the version of the binary
	Version string `json:"version"`
	// Commit is the git commit hash
	Commit string `json:"commit"`
	// Time is the build date
	Time time.Time `json:"time"`
	// GoVersion is the version of the Go toolchain that built the binary
	GoVersion string `json:"go.version"`
}

// Flags contains the build information originated from the ldflags
type Flags struct {
	// Version is the version of the binary
	Version string `validate:"required,semver"`
}

// ParseBuildInfo parses the build information from the flags and the debug build information
func ParseBuildInfo(flags *Flags) (*BuildInfo, error) {
	validate := validator.New()
	isSemverFunc := func(fl validator.FieldLevel) bool {
		return semver.IsValid(fl.Field().String())
	}
	if err := validate.RegisterValidation("semver", isSemverFunc); err != nil {
		return nil, fmt.Errorf("register semver validation: %w", err)
	}
	if err := validate.Struct(flags); err != nil {
		return nil, fmt.Errorf("validate struct: %w", err)
	}
	bi := &BuildInfo{Version: flags.Version}
	dbi, ok := debug.ReadBuildInfo()
	if !ok {
		return bi, nil
	}
	bi.GoVersion = dbi.GoVersion
	for _, s := range dbi.Settings {
		switch s.Key {
		case "vcs.revision":
			bi.Commit = s.Value
		case "vcs.time":
			t, err := time.Parse(time.RFC3339, s.Value)
			if err != nil {
				return nil, fmt.Errorf("parse build date, expected RFC3339 format: %w", err)
			}
			bi.Time = t
		}
	}
	return bi, nil
}
