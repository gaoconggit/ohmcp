package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ide string
	port int
)

var runCmd = &cobra.Command{
	Use:   "run [version]",
	Short: "Run an MCP server for a specific IDE",
	Long: `Start an MCP server instance and configure it for the specified IDE.
If no version is specified, the latest downloaded version will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := "latest"
		if len(args) > 0 {
			version = args[0]
		}
		fmt.Printf("Running MCP server version: %s for IDE: %s on port: %d\n", version, ide, port)
		// TODO: Implement actual server running logic
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	
	runCmd.Flags().StringVarP(&ide, "ide", "i", "cursor", "IDE to configure (cursor, windsurf, cline)")
	runCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the MCP server on")
} 