package main

import (
	"log"
	"net/http"
)

// home writes a byte slice containing text as the response body when navigated to
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// showSnippet is the handler responsible for displaying a specific snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet..."))
}

// createSnippet will allow the user to create a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

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
