package main

import "fmt"

func add(x int, y int) int { //here we can also do (x, y int) like if we have two or more parameters with same data type, we can omit it writing everytime and just keep it in last
	return x + y
}
func main() {
	fmt.Println(add(13, 78))
}
