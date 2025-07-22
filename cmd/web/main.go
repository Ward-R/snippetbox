package main

import (
	"log"
	"net/http"
)

func main() {
	// this starts a new mux(router). sets / pattern to home function
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)                          // Display the home page
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)     // Display a specific snippet
	mux.HandleFunc("GET /snippet/create", snippetCreate)      // Display form for creating new snippet
	mux.HandleFunc("POST /snippet/create", snippetCreatePost) // Save new snippet

	log.Print("starting server on :4000")

	// This starts a new server. Every HTTP request it gets it wills send to the mux
	// to be routed. host:port
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
