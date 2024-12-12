package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) { // Define an http handler function in the standard library form
	respondJSON(w, 200, struct{}{}) // Call the respondJSON function with the response writer, status code, and payload
}
