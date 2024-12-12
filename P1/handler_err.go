package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) { // Define an http handler function in the standard library form
	respondJSON(w, 400, "Something went wrong") // Call the respondJSON function with the response writer, status code, and payload
}
