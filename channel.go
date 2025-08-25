package main

import "fmt"

func main() {
	msg := make(chan string) // This is a unbuffered channel, because it is by default, and it will block until there is another goroutine is ready to recieve
	go func() {              // anonymous function
		msg <- "Hello World"
	}()
	reciever := <-msg
	fmt.Println(reciever)
}
