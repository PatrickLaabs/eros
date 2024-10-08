/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"encoding/json"
	_ "github.com/PatrickLaabs/eros/docs"
	"net/http"
)

//	@Summary		Get API Version
//	@Description	Retrieves the current version of the API, including major, minor, and patch numbers.
//	@Tags			general
//	@Produce		json
//	@Success		200	{object}	api.VersionResponse	"API version response"
//	@Router			/version [get]

// Version returns the current version of the API Server
func Version(w http.ResponseWriter, r *http.Request) {
	version := "0.1.0"
	err := json.NewEncoder(w).Encode(version)
	if err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}
}
