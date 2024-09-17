/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kind

import (
	"encoding/json"
	"log"
	"net/http"
	kind "sigs.k8s.io/kind/pkg/cluster"
)

func GetClusters(w http.ResponseWriter, r *http.Request) {
	p := kind.NewProvider()

	list, err := p.List()
	if err != nil {
		log.Printf("error listing kind clusters: %v", err)
	}

	// Handle successful listing, e.g., return the list as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Printf("error encoding kind clusters: %v", err)
	}
	log.Printf("clusters: %s", list)
}
