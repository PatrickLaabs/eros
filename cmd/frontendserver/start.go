/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package frontendserver

import (
	"github.com/PatrickLaabs/eros/frontend"
	"github.com/spf13/cobra"
)

// startCmd represents the server command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs eros on your bastion node",
	Run: func(cmd *cobra.Command, args []string) {
		frontend.Start()
	},
}
