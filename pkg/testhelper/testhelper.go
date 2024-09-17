package testhelper

import (
	"fmt"
	"net/http"
	"testing"
)

/*
Some helper Functions for the test units
*/

// KubernetesFlavorRequest takes in the kind of flavor we want to process for the kubernetes endpoint
func KubernetesFlavorRequest(flavor string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/kubernetes/%s", flavor), nil)
	return req
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
