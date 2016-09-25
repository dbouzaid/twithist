package hello

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Tests the /hello/ endpoint with a proper name to validate correct results
func TestHelloHandler(t *testing.T) {
	// Create a request with a valid name and no queries
	req, err := http.NewRequest("GET", "/hello/BarackObama", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the response
	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHello)

	// Serve with the response recorder and request
	handler.ServeHTTP(resRec, req)

	// Check if the status of the responses was ok
	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	// Check if what we expected from the output was true
	expected := `Hello BarackObama`
	if resRec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}

func TestHelloNoNameHandler(t *testing.T) {
	// Create a request with no name and no queries
	req, err := http.NewRequest("GET", "/hello/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the response
	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHello)

	// Serve with the response recorder and request
	handler.ServeHTTP(resRec, req)

	// Check if the status of the responses was ok
	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	// Check if what we expected from the output was true
	expected := `Hello `
	if resRec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}

func TestHelloDoubleNameHandler(t *testing.T) {
	// Create a request with a name separated by / and no queries
	req, err := http.NewRequest("GET", "/hello/Barack/Obama", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a recorder to record the response
	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHello)

	// Serve with the response recorder and request
	handler.ServeHTTP(resRec, req)

	// Check if the status of the responses was ok
	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	// Check if what we expected from the output was true
	expected := `Hello Barack`
	if resRec.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}
