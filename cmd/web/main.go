package main

import (
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"basic-go-app.lucassimone.com/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	logger        *slog.Logger
	climbs        *models.Connection
	templateCache map[string]*template.Template
}

func main() {
	// Set the enviroment variables
	setEnvConfig(".env")

	// Instantiate our custom logger
	logger := NewCustomLogger("log.json")

	// Open a database connection pool and check for errors
	db, err := openDB(getEnv("DSN", ""))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// Close the Database connection pool when main ends
	defer db.Close()

	// Create the template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Create our application structure for dependancy injection
	app := &application{
		logger:        logger,
		climbs:        &models.Connection{DB: db},
		templateCache: templateCache,
	}

	logger.Info("Staring server on", "Port", getEnv("DEFAULT_PORT", ":8000"))

	// Get the serve mux from that defines our application routes and serve it
	err = http.ListenAndServe(getEnv("DEFAULT_PORT", ":8000"), app.routes())

	// Exit application on error
	logger.Error(err.Error())
	os.Exit(1)
}

// Opens and returns a sql.DB connection pool to a postgres db for a given dsn
// Pinging the db to see if the connection is successful before returning
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

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
