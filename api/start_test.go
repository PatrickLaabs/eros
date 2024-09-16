/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package api

import (
	"fmt"
	"github.com/PatrickLaabs/eros/api/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
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
