package main

import (
	dct "github.com/PatrickLaabs/eros/structs/dockerclustertemplate"
	"log"
)

/*
Testing interfaces
*/

func main() {
	namespace := "default"
	clustername := "eros-mgmt-cluster"

	_, err := dct.Render(&dct.DockerClusterTemplate{}, clustername, namespace)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
	}
}
