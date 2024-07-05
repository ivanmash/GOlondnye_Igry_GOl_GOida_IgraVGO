package main

import (
	"fmt"
)

func vow(x string) bool {
	switch x {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		return true
	}
	return false
}

func main() {
	var x string
	fmt.Scanln(&x)
	fmt.Println(vow(x))
}
