/*
Copyright © 2026 Ritesh Koushik <riteshkoushik39@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/IAmRiteshKoushik/smriti/internal/runner"
	"github.com/IAmRiteshKoushik/smriti/internal/workspace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run's particular scenarios for MCP testing",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("scenario path required")
		}

		ws, err := workspace.Load(".")
		if err != nil {
			return err
		}

		srv, ok := ws.Servers[ws.Workspace.DefaultServer]
		if !ok {
			return fmt.Errorf("default server not found")
		}

		return runner.RunScenario(context.Background(), srv, args[0])
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
