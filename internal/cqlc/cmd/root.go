package cmd

import "github.com/spf13/cobra"

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "cqlc",
	}
	cmd.AddCommand(NewGenerateCommand())
	return cmd
}
