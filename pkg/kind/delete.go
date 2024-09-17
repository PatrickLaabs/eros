/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kind

import (
	"log"
	"os/exec"
)

func Delete(clustername string) (err error) {
	cmd := exec.Command("kind", "delete", "clusters", clustername)
	err = cmd.Run()
	if err != nil {
		log.Printf("error deleting kind cluster: %v", err)
		return err
	}
	return err
}
