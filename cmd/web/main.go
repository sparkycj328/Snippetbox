package main

import (
	"log"
	"net/http"
)

func main() {
	// declare a new servemux and register the home function as the handler using the `/` as the URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// start a new web server which listens on port 4000 and serves the servemux created above
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
