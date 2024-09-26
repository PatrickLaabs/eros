package main

import (
	"flag"
	"github.com/PatrickLaabs/eros/api/middleware"
	"github.com/PatrickLaabs/eros/api/routes"
	"log"
	"net/http"
)

var (
	backendAddr = flag.String("backendAddr", ":3000", "address to server")
)

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	handler := middleware.Router(mux)

	log.Printf("Running server on port %v", *backendAddr)
	log.Fatal(http.ListenAndServe(*backendAddr, handler))
}
