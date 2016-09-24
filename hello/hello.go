// Contains everything associated with the /hello/ directory of the Web App
package hello

import (
	"bytes"
	"github.com/dbouzaid/twithist/util"
	"net/http"
)

// Used when the user first loads the /hello/ directory
// Gets the name from the endpoint and responses with "Hello [name]"
// where [name] is the name sliced from the path
func LoadHello(w http.ResponseWriter, req *http.Request) {
	// Store the path without the initial "/"
	path := req.URL.Path[1:]
	// Retrieve the name to display from the path
	name := util.GetSlicedPathAt(1, path)

	// Concatenate the main message with the name using a Buffer
	var buffer bytes.Buffer
	buffer.WriteString("Hello ")
	buffer.WriteString(name)

	// Write the response with the output from the buffer
	w.Write(buffer.Bytes())
}
