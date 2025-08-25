package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	ratings := make(chan int, 200) //buffered channel for 200 students to give their ratings
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(studentID int) {
			defer wg.Done()

			sleepTime := rand.Intn(2000)
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)

			rating := rand.Intn(5) + 1
			fmt.Printf("Student %d gave rating %d (after %d ms)\n", studentID, rating, sleepTime)
			ratings <- rating
		}(i)
	}

	go func() {
		wg.Wait()
		close(ratings)
	}()

	sum := 0
	count := 0
	for v := range ratings {
		sum += v
		count++
	}

	avg := float64(sum) / float64(count)
	fmt.Printf("The average rating is %f\n", avg)
}
