package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(writer http.ResponseWriter, request *http.Request) {

	// Modify the header
	writer.Header().Add("Server", "Go")

	// All the template files we need to parse for the html response
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/head.tmpl",
		"./ui/html/partials/header.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/scripts.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Parse the files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	// Write the response with the content of the base template
	err = ts.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

}

func (app *application) pageView(writer http.ResponseWriter, request *http.Request) {

	id, err := strconv.Atoi(request.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}

	//Setting the header content type
	writer.Header().Set("Content-Type", "application/json")
	writer.Write([]byte(`{"name":"Alex"}`))

	//write the response body using another interface
	//fmt.Fprintf(writer, "Display some data with ID %d...", id)
}

func (app *application) pageCreateGet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display a form for creation"))
}

func (app *application) fileGet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display a form for creation"))
	http.ServeFile(writer, request, "./ui/static/img/yellow_boulder.jpg")
}

func (app *application) pageCreatePost(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(http.StatusCreated)

	writer.Write([]byte("Process and save a creation form"))
}
