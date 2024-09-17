/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package middleware

import (
	_ "github.com/PatrickLaabs/eros/docs"
	"net/http"
)

func Router(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allows all origins; adjust as needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass the request to the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
