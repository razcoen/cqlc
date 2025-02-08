package cqlc

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/razcoen/cqlc/internal/buildinfo"
)

// Option is a functional option for configuring the command.
type Option func(*Command)

// Run executes the cqlc command line interface.
func Run(opts ...Option) error {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level:  log.WarnLevel,
		Prefix: "cqlc",
	})

	cmd := &Command{
		Logger: logger,
		BuildFlags: &buildinfo.Flags{
			Version: "v0.0.0-dev",
		},
		Config: &Config{
			DisableOutput: false,
		},
	}

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd.Run()
}

// WithLogger is an option for configuring the logger for the command.
func WithLogger(logger *log.Logger) Option {
	return func(c *Command) {
		c.Logger = logger
	}
}

// WithConfig is an option for configuring the command with a config.
func WithConfig(config *Config) Option {
	return func(c *Command) {
		c.Config = config
	}
}

// WithBuildFlags is an option for configuring the command with build flags.
func WithBuildFlags(flags *buildinfo.Flags) Option {
	return func(c *Command) {
		c.BuildFlags = flags
	}
}
