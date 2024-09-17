/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"fmt"
	_ "github.com/PatrickLaabs/eros/docs"
	"github.com/PatrickLaabs/eros/pkg/kind"
	"log"
	"net/http"
	"strings"
)

// Kubernetes is a Mocking and Test API Implementation
func Kubernetes(w http.ResponseWriter, r *http.Request) {
	flavor := strings.TrimPrefix(r.URL.Path, "/kubernetes/")
	clustername := "eros-bootstrap-cluster"
	// Validate flavor input
	if flavor == "" {
		http.Error(w, "Invalid flavor", http.StatusBadRequest)
		return
	}

	// Handle different flavors
	switch flavor {
	case "local/create":
		fmt.Fprint(w, "Starting local cluster:\n")
		kind.Create(clustername, w, r)
	case "local/delete":
		fmt.Fprint(w, "Deleting local cluster:\n")
		err := kind.Delete(clustername)
		if err != nil {
			log.Printf("error deleting cluster: %v", err)
		}
	case "local/getclusters":
		fmt.Fprint(w, "Getting local clusters:\n")
		kind.GetClusters(w, r)
	case "gcp":
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Starting gcp cluster")
	default:
		http.Error(w, "Unsupported flavor", http.StatusBadRequest)
	}
}
