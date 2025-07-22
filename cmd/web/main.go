package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// command-line flag for HTTP network address
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// initialize structured logger writing to stdout
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// this starts a new mux(router). sets / pattern to home function
	mux := http.NewServeMux()

	// create file server to get files from ./ui/static dir
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use mux to register file server to handle all static paths
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)                          // Display the home page
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)     // Display a specific snippet
	mux.HandleFunc("GET /snippet/create", snippetCreate)      // Display form for creating new snippet
	mux.HandleFunc("POST /snippet/create", snippetCreatePost) // Save new snippet

	logger.Info("starting server", slog.String("addr", *addr))

	// This starts a new server. Every HTTP request it gets it wills send to the mux
	// to be routed. host:port
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
