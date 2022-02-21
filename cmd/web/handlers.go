package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// home writes a byte slice containing text as the response body when navigated to
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// parse template file
	ts, err := template.ParseFiles("../../ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

	data := []byte("Hello from Snippetbox")
	w.Write(data)
}

// showSnippet is the handler responsible for displaying a specific snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Grab the id passed in the request url
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// createSnippet will allow the user to create a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// ensure that createSnippet only executes by using a POST request
	notAllowed := "Method not Allowed"
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, notAllowed, 405)
		return
	}
	data := []byte("Create a new snippet")
	w.Write(data)
}
