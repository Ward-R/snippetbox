package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// This is the home handler function. Hello is the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// This adds a Server: Go header the to response header map.
	w.Header().Add("Server", "Go")

	// Initialize a slice containing paths to our template(HTML) files
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// This reads the template file (HTML) int a template set unless error
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err) // uses serverError() helper
		return
	}

	// Now execute the template set to write the template(HTML) as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err) // uses serverError() helper
	}
}

// snippetView handler
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// snippetCreate handler
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// snippetCreatePost handler
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Send a 201 created status code
	w.WriteHeader(http.StatusCreated)

	// Write response body
	w.Write([]byte("Save a new snippet..."))
}
