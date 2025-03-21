package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"basic-go-app.lucassimone.com/internal/models"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {

	// Modify the header
	writer.Header().Add("Server", "Go")

	app.renderTemplate(writer, request, http.StatusOK, "home.tmpl", templateData{})

}

func (app *application) climbView(writer http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(request.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	// Query the database for the row with the given ID
	climb, err := app.climbs.Get(id)
	// Check for errors returning the appropriate one
	if err != nil {

		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(writer, request)
		} else {
			app.serverError(writer, request, err)
		}

		return
	}

	// Create the tempalte data to pass through
	data := templateData{
		Climb: climb,
	}

	app.renderTemplate(writer, request, http.StatusOK, "view.tmpl", data)

}

func (app *application) climbLatest(writer http.ResponseWriter, request *http.Request) {

	// Get the 10 Latest climbs
	climbs, err := app.climbs.Latest()
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	data := templateData{
		Climbs: climbs,
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/head.tmpl",
		"./ui/html/partials/header.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/scripts.tmpl",
		"./ui/html/pages/latest.tmpl",
	}

	// Parse the templates
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	// Execute the template passing in the climb
	err = ts.ExecuteTemplate(writer, "base", data)
	if err != nil {
		app.serverError(writer, request, err)
	}

}

func (app *application) jsonRequest(writer http.ResponseWriter, request *http.Request) {

	quantity, err := strconv.Atoi(request.PathValue("quantity"))
	if err != nil || quantity < 1 {
		http.NotFound(writer, request)
		return
	}

	if quantity > 100 {
		quantity = 100
	}

	// Get the 10 Latest climbs
	climbs, err := app.climbs.JsonRequest(quantity)

	// Check for errors returning the appropriate one
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(climbs)
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	writer.Write(json)
}

func (app *application) climbCreateGet(writer http.ResponseWriter, request *http.Request) {

	writer.Write([]byte("Display a creation from"))

}

func (app *application) fileGet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display a form for creation"))
	http.ServeFile(writer, request, "./ui/static/img/yellow_boulder.jpg")
}

func (app *application) climbCreatePost(writer http.ResponseWriter, request *http.Request) {

	// Set the header code to the created status... not being used because we redirect to a display page
	//writer.WriteHeader(http.StatusCreated)
	//writer.Write([]byte("Process and save a creation form"))

	// Dummy data
	title, category, grade, setter := "Echo", "boulder", "V5", "john Doe"

	id, err := app.climbs.Insert(title, category, grade, setter)
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	// Redirect the user to the relevant page to view the climb.
	http.Redirect(writer, request, fmt.Sprintf("/view/%d", id), http.StatusSeeOther)
}

// Old climbView function before abstracting to the renderTemplate function
// func (app *application) climbView(writer http.ResponseWriter, request *http.Request) {

// 	id, err := strconv.Atoi(request.PathValue("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(writer, request)
// 		return
// 	}

// 	//Setting the header content type
// 	//writer.Header().Set("Content-Type", "application/json")
// 	//writer.Write([]byte(`{"name":"Alex"}`))

// 	//write the response body using another interface
// 	//fmt.Fprintf(writer, "Display some data with ID %d...", id)

// 	// Query the database for the row with the given ID
// 	climb, err := app.climbs.Get(id)
// 	// Check for errors returning the appropriate one
// 	if err != nil {

// 		if errors.Is(err, models.ErrNoRecord) {
// 			http.NotFound(writer, request)
// 		} else {
// 			app.serverError(writer, request, err)
// 		}

// 		return
// 	}

// 	// Create the tempalte data to pass through
// 	data := templateData{
// 		Climb: climb,
// 	}

// 	// A slice containing the template paths
// 	files := []string{
// 		"./ui/html/base.tmpl",
// 		"./ui/html/partials/head.tmpl",
// 		"./ui/html/partials/header.tmpl",
// 		"./ui/html/partials/footer.tmpl",
// 		"./ui/html/partials/scripts.tmpl",
// 		"./ui/html/pages/view.tmpl",
// 	}

// 	// Parse the templates
// 	ts, err := template.ParseFiles(files...)
// 	if err != nil {
// 		app.serverError(writer, request, err)
// 		return
// 	}

// 	// Execute the template passing in the climb
// 	err = ts.ExecuteTemplate(writer, "base", data)
// 	if err != nil {
// 		app.serverError(writer, request, err)
// 	}
// }
