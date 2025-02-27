package main

import (
	"os"
	"net/http"
	"log/slog"
)

type application struct{
	logger *slog.Logger

}

func main() {
	// Set the enviroment variables
	setEnvConfig(".env")

	// Instantiate our custom logger
	logger := NewCustomLogger("log.json")

	// Create our application structure for dependancy injection
	app := &application{
		logger: logger,
	}

	logger.Info("Staring server on", "Port", getEnv("DEFAULT_PORT", ":8000"))

	// Get the serve mux from that defines our application routes and serve it 
	err := http.ListenAndServe(getEnv("DEFAULT_PORT", ":8000"), app.routes())

	// Exit application on error
	logger.Error(err.Error())
	os.Exit(1)
}
