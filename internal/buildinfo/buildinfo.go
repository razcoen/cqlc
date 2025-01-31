package buildinfo

import (
	"fmt"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/go-playground/validator/v10"
	"golang.org/x/mod/semver"
)

var buildInfo atomic.Pointer[BuildInfo]

// BuildInfo contains information about the build
type BuildInfo struct {
	// Version is the version of the binary
	Version string `json:"version"`
	// Commit is the git commit hash
	Commit string `json:"commit"`
	// Time is the build date
	Time time.Time `json:"time"`
	// GoVersion is the version of the Go toolchain that built the binary
	GoVersion string `json:"go_version"`
}

// Flags contains the build information originated from the ldflags
type Flags struct {
	// Version is the version of the binary
	Version string `validate:"required,semver"`
}

// Store parses the build information and stores it
func Store(flags *Flags) error {
	validate := validator.New()
	if err := validate.RegisterValidation("semver", isSemver); err != nil {
		return fmt.Errorf("register semver validation: %w", err)
	}
	if err := validate.Struct(flags); err != nil {
		return fmt.Errorf("validate struct: %w", err)
	}
	bi := &BuildInfo{
		Version: flags.Version,
	}
	dbi, ok := debug.ReadBuildInfo()
	if !ok {
		buildInfo.Store(bi)
		return nil
	}
	bi.GoVersion = dbi.GoVersion
	for _, s := range dbi.Settings {
		switch s.Key {
		case "vcs.revision":
			bi.Commit = s.Value
		case "vcs.time":
			t, err := time.Parse(time.RFC3339, s.Value)
			if err != nil {
				return fmt.Errorf("parse build date, expected RFC3339 format: %w", err)
			}
			bi.Time = t
		}
	}
	buildInfo.Store(bi)
	return nil
}

// Load returns the build information
func Load() *BuildInfo {
	return buildInfo.Load()
}

func isSemver(fl validator.FieldLevel) bool {
	return semver.IsValid(fl.Field().String())
}
