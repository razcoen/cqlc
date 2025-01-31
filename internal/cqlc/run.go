package cqlc

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/razcoen/cqlc/internal/cqlc/cmd"
	"github.com/spf13/cobra"
)

func Run(flags *buildinfo.Flags) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelWarn,
	}))

	if err := buildinfo.Store(flags); err != nil {
		logger.With("error", err).
			With("flags.version", flags.Version).
			Error("failed to store build info")
		return fmt.Errorf("store build info: %w", err)
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
				*logger = *slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
					AddSource: false,
					Level:     logLevel,
				}))
			}
		},
	}
	rootCmd.PersistentFlags().CountVarP(&options.verbosity, "verbosity", "v", "increase verbosity (-v for info, -vv for debug)")
	rootCmd.AddCommand(cmd.NewGenerateCommand(logger))
	rootCmd.AddCommand(cmd.NewVersionCommand(logger))
	return rootCmd.Execute()
}
