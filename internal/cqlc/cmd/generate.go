package cmd

import (
	"fmt"
	"os"

	"github.com/razcoen/cqlc/pkg/cqlc"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/spf13/cobra"
)

func NewGenerateCommand() *cobra.Command {
	var options struct {
		configPath string
	}

	cmd := &cobra.Command{
		Use: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := os.Open(options.configPath)
			if err != nil {
				return fmt.Errorf("open config file: %w", err)
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()
			cfg, err := config.ParseConfig(f)
			if err != nil {
				panic(err)
			}
			if err := cqlc.Generate(cfg); err != nil {
				panic(err)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&options.configPath, "config", "cqlc.yaml", "generate configuration file")
	return cmd
}
