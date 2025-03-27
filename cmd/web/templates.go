package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"basic-go-app.lucassimone.com/internal/models"
)

// Go Templateing only allows one piece of data to be passed through
// to templates. When we need more data we have to create a struct to
// wrap all the data needed and pass through our templateData struct
type templateData struct {
	Climb       models.Climb
	Climbs      []models.Climb
	CurrentTime time.Time
}

// This function creates a new templateData struct initalizing the global data
//
//	variables and returns it
func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentTime: time.Now(),
	}
}

// Custom template functions can take any amount of parameters but must return only
// one value. The only exception to this is if the function returns an error as the
// second value.
// Takes a time.Time object and returns a string in the format of
// Month Day, Year
func formatMonthDayYear(t time.Time) string {
	return t.Format("January 2, 2006")
}

func currentYear() int {
	return time.Now().Year()
}

// Initalize a template.FuncMap. This acts as a lookup for the custom
// template functions to be used dynamically in our .tmpl files
var functions = template.FuncMap{
	"formatMonthDayYear": formatMonthDayYear,
	"currentYear":  currentYear,
}

// This function simplifies and enhances our handlers in many ways
// 1. It reduces the amount of duplicated code in the handlers pulling all the parsing of templates into one location
// 2. It creates a cache of templates meaning we don't have to parse templates for every request we do it once when starting
// the application and server the same pre parsed templates on each request
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

		// Create a new template with the name of the file
		tmpl := template.New(name)
		// Add the elements of our function map to the templates function map
		tmpl = tmpl.Funcs(functions)
		// Parse the base template. Adding it to the already existing tempalate
		tmpl, err = tmpl.ParseFiles("./ui/html/base.tmpl")
		//The above 3 lines can be executed in a single line
		// Check for any errors
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
