package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/Ward-R/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// This application struct holds application-wide dependencies.
type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// command-line flag for HTTP network address
	addr := flag.String("addr", ":4000", "HTTP network address")
	// command-line flag for MySQL DSN string. ("dsn", "username:password@/snippet...")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// initialize structured logger writing to stdout
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// initialize new instance of application struct
	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", slog.String("addr", *addr))

	// This starts a new server. Every HTTP request it gets it wills send to the mux
	// to be routed. host:port
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// openDB() function wraps sql.Open() returns sql.DB connection pool for given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
