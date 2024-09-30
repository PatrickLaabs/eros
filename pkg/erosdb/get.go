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

func get(erosDbVersion string) (err error) {
	var operatingSystem string

	// Checking OS
	if runtime.GOOS == "darwin" {
		operatingSystem = "macos"
	} else if runtime.GOOS == "linux" {
		operatingSystem = "ubuntu"
	}

	// Retrieves Users home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("cannot read users homedir: %s\n", err)
		return err
	}

	if _, err = os.Stat(filepath.Join(homeDir, ".eros/erosDB-"+operatingSystem+"-latest-"+erosDbVersion)); err != nil {
		if os.IsNotExist(err) {
			//Todo
			log.Printf("no erosDB-%s-latest-%s found: %s\n", operatingSystem, erosDbVersion, err)

			src := "https://github.com/PatrickLaabs/erosDB/releases/download/" + erosDbVersion + "/erosDB-" + operatingSystem + "-latest-" + erosDbVersion
			dst := filepath.Join(homeDir, ".eros/erosDB-"+operatingSystem+"-latest-"+erosDbVersion)

			if err = getter.GetAny(dst, src); err != nil {
				log.Printf("cannot download erosDB-%s-latest-%s: %s\n", operatingSystem, erosDbVersion, err)
				return err
			}

			binPath := filepath.Join(dst, "/erosDB-"+operatingSystem+"-latest-"+erosDbVersion)
			if err = os.Chmod(binPath, 0755); err != nil {
				log.Printf("cannot chmod %s: %s\n", dst, err)
			}
		}
		return err
	}
	return err
}
