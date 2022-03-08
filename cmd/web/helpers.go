package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

// serverError writes an error message and stack trace to the errorLogger than sends a generic
// 500 Internal server error message to the user
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError helper sends a specific status code and corresponding description to the user
// will be used to send errors when a problem with a user's request occurs
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound helper will send a 404 not found response to the user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// render will retrieve the appropriate template set from the cache based on the page name
// as well as write the template to a buffer first to check for errors and then will write an http response
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	var buf *bytes.Buffer

	err := ts.Execute(buf, td)
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}
