package main

import "fmt"

func main() {
	//var ptr *int
	//
	//fmt.Println(ptr) //default value of a pointer if <nil>

	myNumber := 42
	var ptr = &myNumber // reference means &
	fmt.Println(ptr)    // this will give the memory address
	fmt.Println(*ptr)   //this will give the actual value that is stored at that memory address
}
