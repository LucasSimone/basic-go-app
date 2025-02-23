package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/head.tmpl",
		"./ui/html/partials/header.tmpl",
		"./ui/html/partials/footer.tmpl",
		"./ui/html/partials/scripts.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}

}

func pageView(writer http.ResponseWriter, request *http.Request) {

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

func pageCreateGet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display a form for creation"))
}

func fileGet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display a form for creation"))
	http.ServeFile(writer, request, "./ui/static/img/yellow_boulder.jpg")
}

func pageCreatePost(writer http.ResponseWriter, request *http.Request) {

	writer.WriteHeader(http.StatusCreated)

	writer.Write([]byte("Process and save a creation form"))
}
