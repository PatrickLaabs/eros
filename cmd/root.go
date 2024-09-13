/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package cmd

import (
	"github.com/PatrickLaabs/eros/cmd/backendserver"
	"github.com/PatrickLaabs/eros/cmd/frontendserver"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eros",
	Short: "Run Kubernetes in no-time with GitOps onboard",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(backendserver.ServerCmd)
	rootCmd.AddCommand(frontendserver.FrontendCmd)
}
