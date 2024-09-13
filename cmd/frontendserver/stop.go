/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package frontendserver

import (
	"github.com/spf13/cobra"
)

// stopCmd represents the server command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Runs eros on your bastion node",
	Run: func(cmd *cobra.Command, args []string) {
		//frontend.Stop()
	},
}
