package main

import (
	"fmt"
	"net/http" // This package handles all the low-level details of creating the TCP connection, sending the request headers, and receiving the response
)

func main() {
	resp, err := http.Get("https://gemini.google.com/app/0c07fa823b1fb56c")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response header:", resp.Header)
	fmt.Println("Response body:", resp.Body)
}
