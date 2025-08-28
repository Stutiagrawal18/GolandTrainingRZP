package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func worker1(i int) {
	defer wg.Done()
	fmt.Printf("Worker %d starting...\n", i)
	fmt.Printf("Worker %d done.\n", i)
}
func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		worker1(i)
	}
	wg.Wait()
}
