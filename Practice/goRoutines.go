package main

import (
	"fmt"
	"time"
)

func main() {
	greeter("Helllo")
	go greeter("World")
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		fmt.Println(s)
		time.Sleep(2 * time.Second)
	}
}
