/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package api

import (
	"flag"
	"github.com/PatrickLaabs/eros/api/middleware"
	"github.com/PatrickLaabs/eros/api/routes"
	_ "github.com/PatrickLaabs/eros/docs"
	"log"
	"net/http"
)

//	@title			eros swagger API documentation
//	@version		1.0
//	@description	API Backend for the eros platform.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Patrick Laabs
//	@contact.url	http://www.swagger.io/support
//	@contact.email	patrick.laabs@me.com

//	@license.name	GNU GENERAL PUBLIC LICENSE
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/v2

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
