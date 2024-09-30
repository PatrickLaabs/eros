/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package erosdb

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Start starts the erosDB binary after successfully getting it
func Start(erosDbVersion string) (err error) {
	err = get(erosDbVersion)
	if err != nil {
		log.Printf("error getting erosDB binary %s", err)
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("error getting erosDB home dir %s", err)
		return err
	}

	path := filepath.Join(homeDir, ".eros/erosDB-macos-latest-"+erosDbVersion+"/erosDB-macos-latest-"+erosDbVersion)
	cmd := exec.Command(path)

	err = cmd.Run()
	if err != nil {
		log.Printf("Error starting erosDB-macos-latest %s: %s", erosDbVersion, err)
		return err
	}

	return err
}
