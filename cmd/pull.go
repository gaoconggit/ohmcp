package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull [version]",
	Short: "Download an MCP server version",
	Long: `Download a specific version of the MCP server to your local machine.
If no version is specified, the latest stable version will be downloaded.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := "latest"
		if len(args) > 0 {
			version = args[0]
		}
		fmt.Printf("Pulling MCP server version: %s\n", version)
		// TODO: Implement actual download logic
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
} 