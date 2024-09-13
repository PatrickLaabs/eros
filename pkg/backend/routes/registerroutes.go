/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"net/http"
)

var routeMap = map[string]http.HandlerFunc{
	"/version":     Version,
	"/kubernetes/": Kubernetes,
}

//component := hello("John")
//
//	http.Handle("/", templ.Handler(component))

// RegisterRoutes registers all routes with the HTTP server
func RegisterRoutes(mux *http.ServeMux) {
	for path, handler := range routeMap {
		mux.HandleFunc(path, handler)
	}
}
