package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// This is the home handler function. Hello is the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

// snippetView handler
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// snippetCreate handler
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// this starts a new mux(router). sets / pattern to home function
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView) // {id} wildcard
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	// This starts a new server. Every HTTP request it gets it wills send to the mux
	// to be routed. host:port
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
