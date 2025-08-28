package main

import "fmt"

func main() {
	fmt.Println(1)
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")
	fmt.Println(2)
}
