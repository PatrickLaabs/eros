/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package main

import (
	"flag"
	"github.com/PatrickLaabs/eros/api/middleware"
	"github.com/PatrickLaabs/eros/api/routes"
	"github.com/PatrickLaabs/eros/pkg/erosdb"
	"log"
	"net/http"
)

var (
	backendAddr = flag.String("backendAddr", ":3000", "address to server")
)

func main() {
	// starting erosDB in a goroutine along-side API server
	go func() {
		log.Printf("Starting erosDB on Port :3001")
		err := erosdb.Start()
		if err != nil {
			log.Printf("Error starting erosdb: %v", err)
		}
	}()

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	handler := middleware.Router(mux)

	log.Printf("Running server on port %v", *backendAddr)
	log.Fatal(http.ListenAndServe(*backendAddr, handler))
}
