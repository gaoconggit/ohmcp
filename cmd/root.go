package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "omcp",
	Short: "OMCP - A tool for managing MCP servers",
	Long: `OMCP is a command-line tool that helps you manage MCP servers.
It allows you to download, run, and manage different versions of MCP servers.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
} 