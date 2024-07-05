package main

import (
	"fmt"
)

type Rectangle struct {
	width, height int
}

func (a Rectangle) Square() int {
	return a.height * a.width
}

func main() {
	obj := Rectangle{
		width:  10,
		height: 20,
	}
	fmt.Println(obj.Square())
}
