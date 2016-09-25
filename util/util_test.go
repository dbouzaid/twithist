package util

import (
	"testing"
)

func TestGetSlicedPathAt(t *testing.T) {
	path := "/home/hello1//histo gram/"
	pathSliceOne := GetSlicedPathAt(0, path[1:])
	pathSliceTwo := GetSlicedPathAt(1, path[1:])
	pathSliceThree := GetSlicedPathAt(2, path[1:])
	pathSliceFour := GetSlicedPathAt(3, path[1:])


	if pathSliceOne != "home" {
		t.Errorf("GetSlicedPathAt returned wrong value: got %v wanted %v", pathSliceOne, "home")
	}
	if pathSliceTwo != "hello1" {
		t.Errorf("GetSlicedPathAt returned wrong value: got %v wanted %v", pathSliceTwo, "hello1")
	}
	if pathSliceThree != "" {
		t.Errorf("GetSlicedPathAt returned wrong value: got %v wanted %v", pathSliceThree, "")
	}
	if pathSliceFour != "histo gram" {
		t.Errorf("GetSlicedPathAt returned wrong value: got %v wanted %v", pathSliceFour, "histo gram")
	}
}
