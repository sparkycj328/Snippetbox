package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/sparkycj328/Snippetbox/pkg/models"
)

// home writes a byte slice containing text as the response body when navigated to
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	//load in template files into a slice
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

// showSnippet is the handler responsible for displaying a specific snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Grab the id passed in the request url
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// createSnippet will allow the user to create a new snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// ensure that createSnippet only executes by using a POST request
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n-Kobayashi"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
