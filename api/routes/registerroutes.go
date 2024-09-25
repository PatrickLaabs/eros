/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	_ "github.com/PatrickLaabs/eros/docs"
	"net/http"
)

var routeMap = map[string]http.HandlerFunc{
	"/version":     Version,
	"/kubernetes/": Kubernetes,
	"/swagger/":    Swagger,
	"/test":        Test,
}

// RegisterRoutes registers all routes with the HTTP server
func RegisterRoutes(mux *http.ServeMux) {
	for path, handler := range routeMap {
		mux.HandleFunc(path, handler)
	}
}
