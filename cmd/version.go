/*
Copyright © 2026 Ritesh Koushik <riteshkoushik39@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/IAmRiteshKoushik/smriti/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current version: ", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
