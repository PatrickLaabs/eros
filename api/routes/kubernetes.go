/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"fmt"
	"github.com/PatrickLaabs/eros/pkg/kind"
	"net/http"
	"strings"
)

// Kubernetes is a Mocking and Test API Implementation
func Kubernetes(w http.ResponseWriter, r *http.Request) {
	flavor := strings.TrimPrefix(r.URL.Path, "/kubernetes/")

	// Validate flavor input
	if flavor == "" {
		http.Error(w, "Invalid flavor", http.StatusBadRequest)
		return
	}

	// Handle different flavors
	switch flavor {
	case "local":
		fmt.Fprint(w, "Starting local cluster")
		kind.Create()
	case "local/delete":
		fmt.Fprint(w, "Deleting local cluster")
		kind.Delete()
	case "local/getclusters":
		fmt.Fprint(w, "Getting local clusters")
	case "gcp":
		fmt.Fprint(w, "Starting gcp cluster")
	default:
		http.Error(w, "Unsupported flavor", http.StatusBadRequest)
	}

}
