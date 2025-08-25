package main

import (
	"fmt"
	"strings"
	"sync"
)

// countLetters counts frequency of each letter in a word
func countLetters(word string, ch chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done()

	//this map will be created for all the word, later it will go to the channel, and then all the maps from the channel will be merged
	freq := make(map[rune]int) // rune is an alias of int32
	// In the below loop, we are counting the freq of each letter in the word
	for _, r := range strings.ToLower(word) {
		if r >= 'a' && r <= 'z' { // consider only letters
			freq[r]++
		}
	}
	//sending the map created to the channel
	ch <- freq
}

func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	ch := make(chan map[rune]int, len(words)) // buffered channel of size number of words
	//wg is the variable to access add, wait and done of waitgroup
	var wg sync.WaitGroup

	// launch a goroutine for each word
	for _, word := range words {
		//this is telling that we can expect one more co-goroutine call
		wg.Add(1)
		//creating multiple co-goroutines
		go countLetters(word, ch, &wg)
	}

	// close channel when all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// merge results
	finalFreq := make(map[rune]int)
	for partial := range ch {
		for k, v := range partial {
			finalFreq[k] += v
		}
	}

	// print results
	for r := 'a'; r <= 'z'; r++ {
		if count, ok := finalFreq[r]; ok {
			fmt.Printf("%c: %d\n", r, count)
		}
	}
}
