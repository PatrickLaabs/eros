/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"encoding/json"
	"fmt"
	_ "github.com/PatrickLaabs/eros/docs"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Custom ResponseWriter that simulates a failure in the Write method
type failingWriter struct {
	httptest.ResponseRecorder
}

func (fw *failingWriter) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("forced write error")
}

func TestVersion(t *testing.T) {
	t.Run("running versions test for a successful request", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/version", nil)
		response := httptest.NewRecorder()

		Version(response, request)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Decode the response body
		var decodedJson string

		err := json.NewDecoder(response.Body).Decode(&decodedJson)
		if err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		// Verify the response body
		expectedVersion := "0.1.0"
		if decodedJson != expectedVersion {
			t.Errorf("handler returned wrong version: got %v, want %v", decodedJson, expectedVersion)
		}
	})
	//t.Run("running version test for a unsuccessful request", func(t *testing.T) {
	//	// Create a new GET request to the /version endpoint
	//	request := httptest.NewRequest(http.MethodGet, "/version", nil)
	//
	//	// Create a custom failing writer to simulate an encoding error
	//	fw := &failingWriter{*httptest.NewRecorder()}
	//
	//	// Call the Version handler, which will use the failing writer
	//	Version(fw, request)
	//
	//	// Check if the status code is BadRequest due to encoding error
	//	if status := fw.Code; status != http.StatusBadRequest {
	//		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
	//	}
	//
	//	// Check the response body for the error message
	//	expectedBody := "failed to decode json\n"
	//	response := fw.Body.String()
	//
	//	fmt.Printf("response content: %v", response)
	//
	//	if response != expectedBody {
	//		t.Errorf("handler returned unexpected body: got %v, want %v", response, expectedBody)
	//	}
	//})
}
