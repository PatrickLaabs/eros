package routes

import (
	"encoding/json"
	_ "github.com/PatrickLaabs/eros/docs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersion(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/version", nil)
	response := httptest.NewRecorder()

	Version(response, request)

	// Decode the response body
	var decodedJson string

	err := json.NewDecoder(response.Body).Decode(&decodedJson)
	if err != nil {
		t.Errorf("failed to decode json: %v", err)
	}

	assertResponseBody(t, decodedJson, "0.1.0")
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
