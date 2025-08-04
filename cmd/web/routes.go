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

	// Unprotected application routes using the "dynamic" middleware chain.
	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))                      // Display the home page
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView)) // Display a specific snippet
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	// Protected (authenticated-only) application routes, using a new "protected"
	// middleware chain which includes the requireAuthentication middleware.
	protected := dynamic.Append(app.requireAuthentication)
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))      // Display form for creating new snippet
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost)) // Save new snippet
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	// Return the standard middleware chain followed by the servemux.
	return standard.Then(mux)
}
