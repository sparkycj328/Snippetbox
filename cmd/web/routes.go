package main

import "net/http"

// routes declares a new servemux, registers our routes,and returns the servemux back
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	// Create a new file server to server static files and add route
	fileServer := http.FileServer(http.Dir("../../ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return secureHeaders(mux)
}
