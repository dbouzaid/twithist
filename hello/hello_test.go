package hello

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHelloHandler (t *testing.T) {
	req, err := http.NewRequest("GET", "/hello/BarackObama", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHello)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `Hello BarackObama`
	if resRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}

func TestHelloNoNameHandler (t *testing.T) {
	req, err := http.NewRequest("GET", "/hello/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHello)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `Hello `
	if resRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}

func TestHelloDoubleNameHandler (t *testing.T) {
	req, err := http.NewRequest("GET", "/hello/Barack/Obama", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHello)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `Hello Barack`
	if resRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted %v", resRec.Body.String(), expected)
	}
}