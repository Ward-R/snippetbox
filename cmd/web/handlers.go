package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// This is the home handler function. Hello is the response body.
func home(w http.ResponseWriter, r *http.Request) {
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
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Now execute the template set to write the template(HTML) as response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// snippetView handler
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// snippetCreate handler
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// snippetCreatePost handler
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Send a 201 created status code
	w.WriteHeader(http.StatusCreated)

	// Write response body
	w.Write([]byte("Save a new snippet..."))
}
