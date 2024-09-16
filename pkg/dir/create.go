package dir

import (
	"log"
	"os"
)

// Create creates a new directory used for the eros api to save various things.
func Create() {
	//userDir, err := os.UserHomeDir()
	//if err != nil {
	//	log.Printf("error accessing home dir: %v", err)
	//}
	//
	//workingDir := filepath.Join(userDir, ".eros")

	workingDir := Eros()

	// Check if the directory already exists
	if _, err := os.Stat(workingDir); err == nil {
		// Directory already exists, skip creation
		log.Printf("%s already exists: %v", workingDir, err)
	}

	// Create the directory if it doesn't exist
	if err := os.Mkdir(workingDir, 0750); err != nil {
		log.Printf("error creating directory %s: %v", workingDir, err)
	}
}
