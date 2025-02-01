package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/razcoen/cqlc/internal/buildinfo"
	"github.com/spf13/cobra"
)

func NewVersionCommand(logger *slog.Logger, buildInfo *buildinfo.BuildInfo) *cobra.Command {
	var options struct {
		format string
	}

	cmd := &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			info := buildInfo
			var output string
			switch options.format {
			case "text":
				output += "cqlc version " + info.Version + "\n"
				output += "build commit: " + info.Commit + "\n"
				output += "build time: " + info.Time.String() + "\n"
				output += "build go version: " + info.GoVersion + "\n"
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

			_, _ = io.WriteString(cmd.OutOrStdout(), output)
			return nil
		},
	}

	cmd.Flags().StringVar(&options.format, "format", "text", "output format (text, json)")
	return cmd
}
