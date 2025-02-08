package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"

	"github.com/razcoen/cqlc/pkg/cqlc"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/spf13/cobra"
)

func NewGenerateCommand(logger *log.Logger) *cobra.Command {
	var options struct {
		configPath string
	}

	cmd := &cobra.Command{
		Use: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger = logger.With("config", options.configPath)
			logger.Debug("read config file")
			f, err := os.Open(options.configPath)
			if err != nil {
				return fmt.Errorf("open config file: %w", err)
			}
			defer func() {
				if err := f.Close(); err != nil {
					logger.With("error", err).Error("failed to close config file")
				}
			}()
			logger.Debug("parse config file")
			cfg, err := config.ParseConfig(f)
			if err != nil {
				logger.With("error", err).Error("failed to parse config file")
				return nil
			}
			logger.Debug("run cqlc generate")
			if err := cqlc.Generate(cfg, cqlc.WithConfigPath(options.configPath)); err != nil {
				logger.With("error", err).Error("failed during cqlc generate")
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&options.configPath, "config", "cqlc.yaml", "generate configuration file")
	return cmd
}
