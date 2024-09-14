/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package backend

import (
	"encoding/json"
	"fmt"
	"github.com/PatrickLaabs/eros/pkg/backend/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("returns version", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/version", nil)
		response := httptest.NewRecorder()

		routes.Version(response, request)

		// Decode the response body
		var decodedJson string
		//err := json.NewDecoder(response.Body).Decode(&decodedJson)
		err := json.NewDecoder(response.Body).Decode(&decodedJson)
		if err != nil {
			t.Errorf("failed to decode json: %v", err)
		}

		assertResponseBody(t, decodedJson, "0.1.0")
	})
	t.Run("runs kubernetes endpoint using local flavor", func(t *testing.T) {
		request := kubernetesFlavorRequest("local")
		response := httptest.NewRecorder()

		routes.Kubernetes(response, request)
		//assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "Starting local cluster")
	})
	t.Run("runs kubernetes endpoint using gcp flavor", func(t *testing.T) {
		request := kubernetesFlavorRequest("gcp")
		response := httptest.NewRecorder()

		routes.Kubernetes(response, request)
		//assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "Starting gcp cluster")
	})
}

func kubernetesFlavorRequest(flavor string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/kubernetes/%s", flavor), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
