package main

import "fmt"

func summ(nums ...int) int {
	fmt.Println(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
func main() {
	fmt.Println(summ(1, 2, 3), summ(), summ(1, 2, 3, 4, 5))
}
