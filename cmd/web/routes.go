package main

import "net/http"

func (app *application) routes() *http.ServeMux {
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

	return mux
}
