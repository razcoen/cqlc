package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/spf13/cobra"
)

func NewVersionCommand(logger *slog.Logger) *cobra.Command {
	var options struct {
		format string
	}

	cmd := &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			info := buildinfo.Load()
			var output string
			switch options.format {
			case "text":
				output += "version: " + info.Version + "\n"
				output += "commit: " + info.Commit + "\n"
				output += "time: " + info.Time.String() + "\n"
				output += "go version: " + info.GoVersion + "\n"
			case "json":
				b, err := json.MarshalIndent(info, "", "  ")
				if err != nil {
					logger.With("error", err).Error("failed to marshal build info")
					return nil
				}
				output = string(b)
			default:
				return fmt.Errorf("unsupported format: %s", options.format)
			}
			fmt.Print(output)
			return nil
		},
	}

	cmd.Flags().StringVar(&options.format, "format", "text", "output format (text, json)")
	return cmd
}
