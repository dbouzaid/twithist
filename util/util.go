package util

import "strings"

func GetSlicedPathAt(index int, path string) string {
	pathSlice := strings.Split(path, "/")
	name := pathSlice[index]
	return name
}
