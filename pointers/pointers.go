package main

import (
	"fmt"
)

func pointerMagic(pointer *int) {
	*pointer = 7
}

func main() {
	fifty := 50
	one := &fifty
	fmt.Println("Before pointerMagic")
	fmt.Println("-------------------")
	fmt.Println(one)
	fmt.Println(*one)
	pointerMagic(one)
	fmt.Println("After pointerMagic")
	fmt.Println("-------------------")
	fmt.Println(one)
	fmt.Println(*one)
}
