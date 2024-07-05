package main

import (
	"fmt"
)

func prime(x int) {
	var flag bool = false
	for i := 1; i <= x; i += 1 {
		j := i
		for j > 0 {
			if j != 1 && j != i && i%j == 0 {
				flag = true
				break
			}
			j -= 1
		}
		if !flag {
			fmt.Println(i)
		}
		flag = false
	}
}

func main() {
	var x int
	fmt.Scanln(&x)
	prime(x)
}
