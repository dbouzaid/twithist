package home

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHomeHandler (t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHome)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `Try /hello/:name`
	if resRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}

func TestDoubleHomeHandler (t *testing.T) {
	req, err := http.NewRequest("GET", "//", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHome)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `Try /hello/:name`
	if resRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}
