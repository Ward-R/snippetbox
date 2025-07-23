package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// This application struct holds application-wide dependencies.
type application struct {
	logger *slog.Logger
}

func main() {
	// command-line flag for HTTP network address
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// initialize structured logger writing to stdout
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// initialize new instance of application struct
	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", *addr))

	// This starts a new server. Every HTTP request it gets it wills send to the mux
	// to be routed. host:port
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
