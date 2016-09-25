// Contains everything associated with the root directory of the Web App
package home

import "net/http"

// LoadHome is used when the user first loads the root endpoint of the web application.
// Writes out a response to the user to try a different path.
func LoadHome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Try /hello/:name"))
}
