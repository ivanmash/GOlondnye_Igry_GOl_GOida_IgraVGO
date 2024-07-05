package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	var y int
	fmt.Scanln(&y)
	var z int
	fmt.Scanln(&z)
	fmt.Println(max(x, y, z))
}
