package main

import (
	"fmt"
	"time"
)

func worker(i int, done chan bool) {
	fmt.Printf("worker %d starting ", i)
	time.Sleep(3 * time.Second)
	fmt.Printf("worker %d done \n", i)
	done <- true
}
func main() {
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go worker(i, done)
		<-done
	}
}
