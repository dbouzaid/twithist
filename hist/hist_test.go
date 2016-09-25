package hist

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHistHandler (t *testing.T) {
	req, err := http.NewRequest("GET", "/histogram/dhh", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHist)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	notExpected := ``
	if resRec.Body.String() == notExpected {
		t.Errorf("handler returned unexpected body: got %v wanted a JSON Hash", notExpected)
	}
}

func TestHistNoUserHandler (t *testing.T) {
	req, err := http.NewRequest("GET", "/histogram/", nil)
	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	handler := http.HandlerFunc(LoadHist)

	handler.ServeHTTP(resRec, req)

	if status := resRec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v wanted %v", status, http.StatusOK)
	}

	expected := `No user found for specified name`
	if resRec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v wanted a %v", resRec.Body.String(), expected)
	}
}
