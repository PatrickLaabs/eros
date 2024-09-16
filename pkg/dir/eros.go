package dir

import (
	"log"
	"os"
	"path/filepath"
)

func Eros() (workingDir string) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("error accessing home dir: %v", err)
	}

	workingDir = filepath.Join(userDir, ".eros")
	return workingDir
}
