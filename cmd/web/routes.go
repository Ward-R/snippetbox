package main

import (
	"net/http"

	"github.com/Ward-R/snippetbox/ui"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// this starts a new mux(router). sets / pattern to home function
	mux := http.NewServeMux()

	// Use the http.FileServerFS() function to create an HTTP handler which
	// serves the embedded files in ui.Files. It's important to note that our
	// static files are contained in the "static" folder of the ui.Files
	// embedded filesystem. So, for example, our CSS stylesheet is located at
	// "static/css/main.css". This means that we no longer need to strip the
	// prefix from the request URL -- any requests that start with /static/ can
	// just be passed directly to the file server and the corresponding static
	// file will be served (so long as it exists).

	// Use mux to register file server to handle all static paths
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	// Add a new GET /ping route.
	mux.HandleFunc("GET /ping", ping)

	// Unprotected application routes using the "dynamic" middleware chain.
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

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
