package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int32
	counter2 := 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter2++
				counter.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final counter", counter.Load(), counter2)
}
