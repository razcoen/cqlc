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
			cfg, err := cqlc.ReadConfig("./cqlc.yaml")
			if err != nil {
				panic(err)
			}
			for _, c := range cfg.CQL {
				if err := g.Generate(c); err != nil {
					panic(err)
				}
			}
			return nil
		},
	}
	return cmd
}
