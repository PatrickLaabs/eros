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
func Start() (err error) {
	version := "0.0.12"
	
	err = get(version)
	if err != nil {
		log.Printf("error getting erosDB binary %s", err)
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("error getting erosDB home dir %s", err)
		return err
	}

	path := filepath.Join(homeDir, ".eros/erosDB-macos-latest/erosDB-macos-latest")
	cmd := exec.Command(path)

	err = cmd.Run()
	if err != nil {
		log.Printf("Error starting erosDB-macos-latest: %s", err)
		return err
	}

	return err
}
