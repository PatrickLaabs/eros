/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

// Version returns the current version of the API Server
func Version(w http.ResponseWriter, r *http.Request) {
	version := "0.1.0"
	err := json.NewEncoder(w).Encode(version)
	if err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
