package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ward-R/snippetbox/internal/models"
)

// This is the home handler function. Hello is the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// This adds a Server: Go header the to response header map.
	w.Header().Add("Server", "Go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Use render helper
	app.render(w, r, http.StatusOK, "home.tmpl", templateData{
		Snippets: snippets,
	})
}

// snippetView handler
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use SnippetModel's Get() method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// Return a 404 Not Found Response.
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Use render helper
	app.render(w, r, http.StatusOK, "view.tmpl", templateData{
		Snippet: snippet,
	})
}

// snippetCreate handler
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// snippetCreatePost handler
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Create seed data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Redirect the user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
