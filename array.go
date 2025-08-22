package main

import "fmt"

func main() {
	fmt.Println("Array")
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
