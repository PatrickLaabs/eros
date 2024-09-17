/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	_ "github.com/PatrickLaabs/eros/docs"
	"github.com/swaggo/http-swagger"
	"net/http"
)

func Swagger(w http.ResponseWriter, r *http.Request) {
	httpSwagger.WrapHandler(w, r)
}
