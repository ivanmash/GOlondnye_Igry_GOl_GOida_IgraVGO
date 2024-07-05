package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	var x int
	fmt.Scanln(&x)
	var y int
	fmt.Scanln(&y)
	var sum int = add(x, y)
	fmt.Println(sum)
}
