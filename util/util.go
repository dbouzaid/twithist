// Package util contains code that can be used by multiple packages associated with the
// web application
package util

import "strings"

// GetSlicedPathAt takes a URL path and splits it by /. It then uses the given
// index to return the part of the URL path wanted.
func GetSlicedPathAt(index int, path string) string {
	pathSlice := strings.Split(path, "/")
	name := pathSlice[index]
	return name
}
