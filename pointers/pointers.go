package main

import (
	"fmt"
)

func pointerMagic(pointer *int) {
	*pointer = 7
}

func main() {
	fifty := 50
	seven := &fifty
	fmt.Println("Before pointerMagic")
	fmt.Println("-------------------")
	fmt.Println(seven)
	fmt.Println(*seven)
	pointerMagic(seven)
	fmt.Println("After pointerMagic")
	fmt.Println("-------------------")
	fmt.Println(seven)
	fmt.Println(*seven)
}
