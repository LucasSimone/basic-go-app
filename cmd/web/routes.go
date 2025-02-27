package main

import (
	"net/http"
)

// The routes function returns a ServeMux which defines the various routes for our application
func(app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Default go fileserver handler to handle the static files
	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	//mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Custom Neutered file server. This file server will only allow the serving of
	// specified files. If a directory path is given it will return a 404
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(getEnv("STATIC_DIRECTORY", "/static"))})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Handlers execute each within their own goroutine. This means for mutiple requests code
	// will be called at the same time leading to race conditions on handlers that access the
	// same resources

	// Prefixing the route patterns with the required HTTP method with
	// restrict the routes to only act on their respective requests.

	// {$} restricts any sub tree paths otherwise home would be served on every page
	// that doesn not fit a more specific route
	mux.HandleFunc("GET /{$}", app.home)

	mux.HandleFunc("GET /view/{id}", app.pageView)
	mux.HandleFunc("GET /create", app.pageCreateGet)
	mux.HandleFunc("GET /file", app.fileGet)

	// Seperate POST request route for the create route
	mux.HandleFunc("POST /create", app.pageCreatePost)

	return mux
}