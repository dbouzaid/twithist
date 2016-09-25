package hist

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests the /histogram/ endpoint with a valid username to see if we get
// a valid JSON hash
func TestHistHandler(t *testing.T) {
	// Create a request with a valid twitter username and no queries
	req, err := http.NewRequest("GET", "/histogram/dhh", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the response
	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHist)

	// Serve with the response recorder and request
	handler.ServeHTTP(resRec, req)

	// Check if the status of the responses was ok
	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	// Check if what we received a JSON hash instead of nothing
	notExpected := ``
	if resRec.Body.String() == notExpected {
		t.Errorf("Handler returned unexpected body: got %v wanted a JSON Hash", notExpected)
	}
}

// Tests the /histogram/ endpoint with no name to see if we get the right error
// message
func TestHistNoUserHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/histogram/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the response
	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHist)

	// Serve with the response recorder and request
	handler.ServeHTTP(resRec, req)

	// Check if the status of the responses was ok
	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	// Check if what we expected from the output was true
	expected := `No user found for specified name`
	if resRec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted a %v", resRec.Body.String(), expected)
	}
}
