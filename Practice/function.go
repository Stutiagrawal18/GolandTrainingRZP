package main

import "fmt"

func add(x int, y int) int { //here we can also do (x, y int) like if we have two or more parameters with same data type, we can omit it writing everytime and just keep it in last
	return x + y
}

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}
func main() {
	//fmt.Println(add(13, 78))
	fmt.Println(concat("hello", "world"))
}

//naked return
//A return statement without arguements returns the named return values. This is known as naked return
