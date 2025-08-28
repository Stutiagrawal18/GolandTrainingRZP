package main

import (
	"fmt"
	"net/http"
)

// The handler function for the root URL "/"
// It takes a http.ResponseWriter to write the response, and a *http.Request
// to read from the request.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// The Fprintf function writes a formatted string to the http.ResponseWriter.
	// We are sending a simple "Hello, World!" message.
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// http.HandleFunc is used to register a handler function with a URL path.
	// When a request comes in for "/", it will be handled by helloHandler.
	http.HandleFunc("/hello", helloHandler)

	// http.ListenAndServe starts the server. It takes the address to listen on
	// (":8080" means all network interfaces on port 8080) and a handler.
	// Since we've already registered handlers, we pass nil here.
	// It returns an error, which we handle by panicking if it fails.
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// If ListenAndServe returns an error, it's usually unrecoverable.
		// So we use panic to stop the program.
		panic(err)
	}
}
