package main

import (
	"log"
	"net/http"
)

func main() {
	setEnvConfig()

	mux := http.NewServeMux()
	// Read in the .env

	// Handlers execute each within their own goroutine. This means for mutiple requests code
	// will be called at the same time leading to race conditions on handlers that access the
	// same resources

	// {$} prevents the home route from being served on all sub paths
	// mux.HandleFunc("/{$}", home)

	// Prefixing the route patterns with the required HTTP method with
	// restrict the routes to only act on their respective requests.
	mux.HandleFunc("GET /{$}", home)

	mux.HandleFunc("GET /view/{id}", pageView)
	mux.HandleFunc("GET /create", pageCreateGet)
	mux.HandleFunc("GET /file", fileGet)

	// Seperate POST request route for the create route
	mux.HandleFunc("POST /create", pageCreatePost)

	// Default go fileserver handler to handle the static files
	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	//mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Custom Neutered file server. This file server will only allow the serving of
	// specified files. If a directory path is given it will return a 404
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(getEnv("STATIC_DIRECTORY", "/static"))})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Staring server on", getEnv("DEFAULT_PORT", ":8000"))

	err := http.ListenAndServe(getEnv("DEFAULT_PORT", ":8000"), mux)
	log.Fatal(err)
}
