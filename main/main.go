package main

import (
	"net/http"

	"github.com/dbouzaid/twithist/hello"
	"github.com/dbouzaid/twithist/hist"
	"github.com/dbouzaid/twithist/home"
)

func main() {
	// Based on where the user has pointed to in their web browser,
	// load the corresponding function
	http.HandleFunc("/", home.LoadHome)
	http.HandleFunc("/hello/", hello.LoadHello)
	http.HandleFunc("/histogram/", hist.LoadHist)

	// Using port 1010 as 10 is my lucky number :)
	http.ListenAndServe(":1010", nil)
}
