/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mock the kind package and its GetClusters function
var clusterMock = struct {
	GetClusters func(w http.ResponseWriter, r *http.Request)
}{
	GetClusters: func(w http.ResponseWriter, r *http.Request) {
		// Simulate response from the real GetClusters function
		fmt.Fprint(w, "Mocked clusters data")
	},
}

func TestKubernetes(t *testing.T) {
	t.Run("runs on no given flavor, should error", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/kubernetes/", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Call the Kubernetes function with the ResponseRecorder and the request
		handler := http.HandlerFunc(Kubernetes)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

		// Check the response body
		expected := "Invalid flavor"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
	t.Run("runs kubernetes endpoint using local flavor", func(t *testing.T) {

		//// ToDo: Needs to be adjusted to the new way we do things
		//request := testhelper.KubernetesFlavorRequest("local/create")
		//response := httptest.NewRecorder()
		//
		//// ToDo: currently our test creates a cluster. We should mock this
		//Kubernetes(response, request)
		////assertStatus(t, response.Code, http.StatusOK)
		//testhelper.AssertResponseBody(t, response.Body.String(), "Starting local cluster")
	})
	t.Run("running test for getclusters endpoint", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/kubernetes/local/getclusters", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Kubernetes)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expectedBody := "Getting local clusters:\n[\"eros-bootstrap-cluster\"]"
		if strings.TrimSpace(rr.Body.String()) != expectedBody {
			t.Errorf("Handler returned unexpected body: got: %v want: %v", rr.Body.String(), expectedBody)
		}
	})
	t.Run("run tests on gcp flavor", func(t *testing.T) {
		// Create a new request with the "gcp" flavor (e.g., "/kubernetes/gcp")
		req, err := http.NewRequest("GET", "/kubernetes/gcp", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Call the Kubernetes function with the ResponseRecorder and the request
		handler := http.HandlerFunc(Kubernetes)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check the response body
		expected := "Starting gcp cluster"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
	t.Run("runs kubernetes endpoint with the default case", func(t *testing.T) {
		// Create a new request with an unsupported flavor (e.g., "/kubernetes/unsupported")
		req, err := http.NewRequest("GET", "/kubernetes/unsupported", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(Kubernetes)
		handler.ServeHTTP(rr, req)

		// Check the response status code
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

		// Check the response body
		expected := "Unsupported flavor"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}
