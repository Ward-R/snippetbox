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

	// Create a new middleware chain containing the middleware specific to our
	// dynamic application routes. For now, this chain will only contain the
	// LoadAndSave session middleware but we'll add more to it later.
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))                          // Display the home page
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))     // Display a specific snippet
	mux.Handle("GET /snippet/create", dynamic.ThenFunc(app.snippetCreate))      // Display form for creating new snippet
	mux.Handle("POST /snippet/create", dynamic.ThenFunc(app.snippetCreatePost)) // Save new snippet

	// The middleware chain previously looked like this:
	// return app.recoverPanic(app.logRequest(commonHeaders(mux)))

	// It has been changed to look cleaner with justinas/alice package
	// Middleware chain conaining the standard middleware
	// this will be used for every request our application receives
	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	// Return the standard middleware chain followed by the servemux.
	return standard.Then(mux)
}
