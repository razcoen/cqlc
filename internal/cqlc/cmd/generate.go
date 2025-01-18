package cmd

import (
	"github.com/razcoen/cqlc/pkg/cqlc"
	"github.com/spf13/cobra"
)

func NewGenerateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "generate",
		RunE: func(cmd *cobra.Command, args []string) error {
			g, err := cqlc.NewGenerator()
			if err != nil {
				panic(err)
			}
			cfg, err := cqlc.ReadConfig("cqlc.yaml")
			if err != nil {
				panic(err)
			}
			if err := g.Generate(cfg); err != nil {
				panic(err)
			}
			return nil
		},
	}
	return cmd
}
