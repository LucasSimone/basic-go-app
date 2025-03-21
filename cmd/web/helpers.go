package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// The serverError function writes a log entry with relevant data (method, uri), then replies to
// the request with a 500 Internal server error
func (app *application) serverError(writer http.ResponseWriter, request *http.Request, err error) {
	method := request.Method
	uri := request.URL.RequestURI()

	// Include the trace variable in the log to get the stack trace,
	// outlining the execution path of the application
	//trace := string(debug.Stack())

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The client function replies to the request with a specified status code
func (app *application) clientError(writer http.ResponseWriter, status int) {
	http.Error(writer, http.StatusText(status), status)
}

// renderTemplate is used to execute templates and check for errors before sending them to the http.ResponseWriter
// This reduces the duplicated code in the handlers and simplifies serveing templates
func (app *application) renderTemplate(writer http.ResponseWriter, request *http.Request, status int, page string, data templateData) {

	// Get the template
	tmpl, found := app.templateCache[page]
	// Check if the template was found returning a new error if not
	if !found {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(writer, request, err)
		return
	}

	// Create a writer buffer to write the template execution to instead of the http.ResponseWriter
	writerBuffer := new(bytes.Buffer)

	// Execute the template and check for errors
	err := tmpl.ExecuteTemplate(writerBuffer, "base", data)
	if err != nil {
		app.serverError(writer, request, err)
		return
	}

	// Add the Http status code
	writer.WriteHeader(status)

	// Writer the buffer to the http.ResponseWriter
	writerBuffer.WriteTo(writer)

}
