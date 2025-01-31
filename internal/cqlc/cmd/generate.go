package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/razcoen/cqlc/pkg/cqlc"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/spf13/cobra"
)

func NewGenerateCommand(logger *slog.Logger) *cobra.Command {
	var options struct {
		configPath string
	}

	cmd := &cobra.Command{
		Use: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger = logger.With("config", options.configPath)
			f, err := os.Open(options.configPath)
			if err != nil {
				return fmt.Errorf("open config file: %w", err)
			}
			defer func() {
				if err := f.Close(); err != nil {
					logger.With("error", err).Error("failed to close config file")
				}
			}()
			cfg, err := config.ParseConfig(f)
			if err != nil {
				logger.With("error", err).Error("failed to parse config file")
				return nil
			}
			if err := cqlc.Generate(cfg); err != nil {
				logger.With("error", err).Error("failed during cqlc generate")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&options.configPath, "config", "cqlc.yaml", "generate configuration file")
	return cmd
}
