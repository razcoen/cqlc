package cmd

import (
	"fmt"
	"github.com/razcoen/cqlc/pkg/cqlc"
	"github.com/razcoen/cqlc/pkg/cqlc/config"
	"github.com/spf13/cobra"
	"os"
)

func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			pf := cmd.Flag("config")
			cfgpath := pf.Value.String()
			f, err := os.Open(cfgpath)
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
	// TODO: Bind flags to viper config
	cmd.Flags().String("config", "cqlc.yaml", "generate configuration file")
	return cmd
}
