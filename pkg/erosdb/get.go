/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package erosdb

import (
	"github.com/hashicorp/go-getter"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func get(version string) (err error) {
	var operatingSystem string

	if runtime.GOOS == "darwin" {
		operatingSystem = "macos"
		// https://github.com/PatrickLaabs/erosDB/releases/download/0.0.11/erosDB-macos-latest
	} else if runtime.GOOS == "linux" {
		operatingSystem = "ubuntu"
		// https://github.com/PatrickLaabs/erosDB/releases/download/0.0.11/erosDB-ubuntu-latest
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("cannot read users homedir: %s\n", err)
		return err
	}

	if _, err = os.Stat(filepath.Join(homeDir, ".eros/erosDB-"+operatingSystem+"-latest")); err != nil {
		if os.IsNotExist(err) {
			log.Printf("no erosDB-macos-latest found\n")

			src := "https://github.com/PatrickLaabs/erosDB/releases/download/" + version + "/erosDB-" + operatingSystem + "-latest"
			dst := filepath.Join(homeDir, ".eros/erosDB-macos-latest")

			if err = getter.GetAny(dst, src); err != nil {
				log.Printf("cannot download erosDB-macos-latest: %s\n", err)
				return err
			}
		}
		return err
	}
	return err
}
