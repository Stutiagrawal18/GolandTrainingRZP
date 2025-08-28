package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}                 //we mandatory have to initialise this here only else it gives an error
	fmt.Printf("type of slice is %T\n", slice) // type is []int
	fmt.Println(slice)                         // [1, 2, 3] ; this print the entire slice

	//to add data manually, we use the append method

	slice = append(slice, 4)
	fmt.Println(slice)
}
