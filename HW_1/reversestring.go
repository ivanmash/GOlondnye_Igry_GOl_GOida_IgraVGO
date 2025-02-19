package main

import (
	"fmt"
)

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {
	var s string
	fmt.Scanln(&s)
	fmt.Println(Reverse(s))
}
