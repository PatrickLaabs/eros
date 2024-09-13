/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package frontendserver

import (
	"fmt"
	"github.com/spf13/cobra"
)

// FrontendCmd represents the server command
var FrontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "Runs eros on your bastion node",
	Long: `Runs eros on your bastion node` +
		`and will take care of your management cluster using Cluster-API`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	FrontendCmd.AddCommand(startCmd)
	FrontendCmd.AddCommand(stopCmd)
}
