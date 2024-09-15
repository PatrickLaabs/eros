/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package middleware

import (
	"github.com/PatrickLaabs/eros/api/routes"
	"net/http"
)

func Router(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the URL path
		path := r.URL.Path

		// Implement routing logic based on the URL path
		switch path {
		case "/version":
			routes.Version(w, r)
		case "/kubernetes":
			routes.Kubernetes(w, r)
		default:
			// Handle 404 Not Found
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		// Pass the request to the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
