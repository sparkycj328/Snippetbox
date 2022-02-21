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
	data := []byte("Hello from Snippetbox")
	w.Write(data)
}

// showSnippet is the handler responsible for displaying a specific snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	data := []byte("Display a specific snippet...")
	w.Write(data)
}

// createSnippet will allow the user to create a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// ensure that createSnippet only executes by using a POST request
	notAllowed := []byte("Method Not Allowed")
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write(notAllowed)
		return
	}
	data := []byte("Create a new snippet")
	w.Write(data)
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
