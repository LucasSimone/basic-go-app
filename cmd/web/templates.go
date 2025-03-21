package main

import (
	"html/template"
	"path/filepath"

	"basic-go-app.lucassimone.com/internal/models"
)

// Go Templateing only allows one piece of data to be passed through
// to templates. When we need more data we have to create a struct to
// wrap all the data needed and pass through our templateData struct
type templateData struct {
	Climb  models.Climb
	Climbs []models.Climb
}

// This function simplifies and enhances our handlers in many ways
// 1. It reduces the amount of duplicated code in the handlers pulling all the parsing of templates into one location
// 2. It creates a cache of templates meaning we don't have to parse templates for every request we do it once when starting
// the application and server the same pre parsed templates on each request
// 3.
func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a map of Templates to act as the cache
	cache := map[string]*template.Template{}

	// Get a slice of filepaths strings for all files in our /ui/html/pages folder
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the name/page of the file for each filepath
		name := filepath.Base(page)

		// Parse the base file
		tmpl, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Parse the partials. Adding them into the already existing tempalate
		tmpl, err = tmpl.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Parse the page. Adding it to the already existing tempalate
		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the created template to the cache
		cache[name] = tmpl
	}

	return cache, nil
}
