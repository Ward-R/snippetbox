package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// this starts a new mux(router). sets / pattern to home function
	mux := http.NewServeMux()

	// create file server to get files from ./ui/static dir
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use mux to register file server to handle all static paths
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)                          // Display the home page
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)     // Display a specific snippet
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)      // Display form for creating new snippet
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost) // Save new snippet

	// The middleware chain previously looked like this:
	// return app.recoverPanic(app.logRequest(commonHeaders(mux)))

	// It has been changed to look cleaner with justinas/alice package
	// Middleware chain conaining the standard middleware
	// this will be used for every request our application receives
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	// Return the standard middleware chain followed by the servemux.
	return standard.Then(mux)
}
