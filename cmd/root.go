/*
Copyright © 2026 Ritesh Koushik <riteshkoushik39@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/IAmRiteshKoushik/smriti/internal/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "smriti",
	Short:   "CLI and CI runner for testing MCP servers",
	Version: version.Version,
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
