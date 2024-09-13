/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package frontend

import (
	"flag"
	"github.com/PatrickLaabs/eros/views"
	"github.com/a-h/templ"
	"log"
	"net/http"
)

/*
Frontend Application.
Consuming the backend api to created and manage kubernetes clusters.

This will always be started, when running the backend server.
*/

var (
	frontendAddr = flag.String("frontendAddr", ":8080", "address to server")
)

func Start() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.Handle("/", templ.Handler(views.LandingPage()))

	log.Printf("starting frontend on port %v", *frontendAddr)
	log.Fatal(http.ListenAndServe(*frontendAddr, nil))
}
