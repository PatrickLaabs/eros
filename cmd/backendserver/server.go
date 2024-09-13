/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package backendserver

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ServerCmd represents the server command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs eros on your bastion node",
	Long: `Runs eros on your bastion node` +
		`and will take care of your management cluster using Cluster-API`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	ServerCmd.AddCommand(startCmd)
	ServerCmd.AddCommand(stopCmd)
}
