package main

import (
	"fmt"
)

func odd_even(x int) bool {
	return x%2 == 0
}

func main() {
	var x int
	fmt.Scanln(&x)
	if odd_even(x) {
		fmt.Println("Чётное.")
	} else {
		fmt.Println("Нечётное.")
	}
}
