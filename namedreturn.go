package main

import "fmt"

func sum(sum int) (x, y int) {
	x = sum + 10
	y = sum - 10
	return
}

func main() {
	fmt.Println(sum(50))
}
