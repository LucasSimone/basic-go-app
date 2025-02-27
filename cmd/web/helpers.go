package main

import "net/http"

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