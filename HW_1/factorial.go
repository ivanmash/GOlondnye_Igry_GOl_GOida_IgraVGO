package main

import (
	"fmt"
)

func factorial(x int) int {
	total := 1
	for x > 0 {
		total = total * x
		x -= 1
	}
	return total
}

func main() {
	var x int
	fmt.Scanln(&x)
	fmt.Println(factorial(x))
}
