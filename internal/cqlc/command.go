package cqlc

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/razcoen/cqlc/internal/cqlc/cmd"
	"github.com/spf13/cobra"
)

// Command is the main cqlc command.
type Command struct {
	// Logger is the logger for the command.
	Logger *slog.Logger
	// BuildFlags are the build flags for the command.
	BuildFlags *buildinfo.Flags
	// Config is the internal configuration for the command.
	Config *Config
}

// Config is the non user facing configuration for the cqlc command.
type Config struct {
	// DisableOutput disables output to stdout ann stderr.
	// This is mainly used for testing to avoid printing to the terminal.
	DisableOutput bool
}

// Run executes the cqlc command line interface.
func (c *Command) Run() error {
	buildInfo, err := buildinfo.ParseBuildInfo(c.BuildFlags)
	if err != nil {
		return fmt.Errorf("parse build info: %w", err)
	}

	var options struct {
		verbosity int
	}

	rootCmd := &cobra.Command{
		Use: "cqlc",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logLevel := slog.LevelInfo
			if options.verbosity >= 2 {
				logLevel = slog.LevelDebug
			}
			if options.verbosity > 0 {
				*c.Logger = *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
					AddSource: false,
					Level:     logLevel,
				}))
			}
		},
	}

	if c.Config.DisableOutput {
		rootCmd.SetUsageFunc(func(c *cobra.Command) error { return nil })
	}

	rootCmd.PersistentFlags().CountVarP(&options.verbosity, "verbosity", "v", "increase verbosity (-v for info, -vv for debug)")
	rootCmd.AddCommand(cmd.NewGenerateCommand(c.Logger))
	rootCmd.AddCommand(cmd.NewVersionCommand(c.Logger, buildInfo))
	return rootCmd.Execute()
}
