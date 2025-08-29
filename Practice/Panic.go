package main

import (
	"fmt"
	"sync"
	"time"
)

// The worker function simulates a task that will panic.
// It includes a defer function with a recover call.
func worker2(id int, wg *sync.WaitGroup) {
	// The defer statement schedules the function call to be executed
	// just before the surrounding function returns. This is crucial for recover().
	defer func() {
		// Use recover() to capture the panic. It returns the value passed to panic().
		if r := recover(); r != nil {
			fmt.Printf("Goroutine %d has been recovered from a panic: %v\n", id, r)
		}
		// Acknowledge that this goroutine is done.
		wg.Done()
	}()

	// Print a message to show the goroutine has started.
	fmt.Printf("Goroutine %d started.\n", id)

	// Simulate a random delay to make the output a bit more interesting.
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)

	// This is the line that will cause the panic.
	// We're panicking with a simple string value.
	panic(fmt.Sprintf("Something went wrong in goroutine %d!", id))
}

func main() {
	// Use a WaitGroup to wait for all goroutines to finish.
	// This ensures the main function doesn't exit before the workers complete.
	var wg sync.WaitGroup

	// Set the number of goroutines we want to launch.
	const numGoroutines = 20

	// Add the number of goroutines to the WaitGroup.
	wg.Add(numGoroutines)

	// Loop to launch 20 goroutines.
	for i := 1; i <= numGoroutines; i++ {
		// Launch a new goroutine for each worker.
		// We pass the id and the WaitGroup pointer.
		go worker2(i, &wg)
	}

	// Wait for all the goroutines in the WaitGroup to complete.
	fmt.Println("Waiting for all goroutines to panic and recover...")
	wg.Wait()

	// This line will be reached because all panics were recovered.
	fmt.Println("All goroutines have finished and been recovered. Main program continues.")
}
