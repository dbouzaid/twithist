package home

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests the root endpoint to validate correct results
func TestHomeHandler(t *testing.T) {
	// Create a request with the root path and no queries
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the response
	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHome)

	// Serve with the response recorder and request
	handler.ServeHTTP(resRec, req)

	// Check if the status of the responses was ok
	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	// Check if what we expected from the output was true
	expected := `Try /hello/:name`
	if resRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}
