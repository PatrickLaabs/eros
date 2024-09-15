/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package backendserver

import (
	"fmt"
	"github.com/PatrickLaabs/eros/api"
	"github.com/spf13/cobra"
)

// startCmd represents the server command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs eros on your bastion node",
	Run: func(cmd *cobra.Command, args []string) {

		/* ToDos:
		1. Set up Kind
		*/

		//kind.Create()
		fmt.Println("server-backend start called")
		api.Start()
	},
}

//func init() {
//	serverCmd.AddCommand(startCmd)
//}
