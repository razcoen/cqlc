package cmd

import "github.com/spf13/cobra"

func NewGenerateCommand() *cobra.Command {
  cmd := &cobra.Command{
    Use: "generate",
  }
  return cmd
}
