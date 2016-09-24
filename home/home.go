// Contains everything associated with the root directory of the Web App
package home

import "net/http"

// Used when the user first loads the root page of the web application.
// Writes out a message to the user to try a different path.
func LoadHome(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Try /hello/:name"))
}
