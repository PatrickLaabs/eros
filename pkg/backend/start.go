/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package backend

import (
	"flag"
	"github.com/PatrickLaabs/eros/pkg/backend/middleware"
	"github.com/PatrickLaabs/eros/pkg/backend/routes"
	"log"
	"net/http"
)

/*
Backend api server.
Provides APIs for the frontend to consume, as creating and managing kubernetes clusters.
*/

var (
	backendAddr = flag.String("backendAddr", ":3000", "address to server")
)

func Start() {
	mux := http.NewServeMux()
	mux.Handle("/", middleware.Router(mux))

	routes.RegisterRoutes(mux)

	log.Printf("Running server on port %v", *backendAddr)
	log.Fatal(http.ListenAndServe(*backendAddr, mux))
}
