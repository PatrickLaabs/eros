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

func Create(clustername string, w http.ResponseWriter, r *http.Request) {
	p := kind.NewProvider()
	if err := p.Create(
		clustername,
		kind.CreateWithRawConfig(configgen()),
	); err != nil {
		log.Printf("error creating kind cluster: %v", err)
	}

	w.Header().Set("Content-Type", "application/yaml")
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Printf("error encoding kind cluster: %v", err)
	}
	//log.Printf("created kind cluster: %s", p)
}
