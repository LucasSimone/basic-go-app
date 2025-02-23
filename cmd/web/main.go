package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

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
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Print("Staring server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
