/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package backendserver

import (
	"fmt"
	"github.com/spf13/cobra"
)

// stopCmd represents the server command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Runs eros on your bastion node",
	Run: func(cmd *cobra.Command, args []string) {

		/* ToDos:
		1. Set up Kind
		*/

		//kind.Create()
		fmt.Println("server stop called")
	},
}

//func init() {
//	serverCmd.AddCommand(startCmd)
//}
