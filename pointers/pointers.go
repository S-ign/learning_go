package main

import (
	"fmt"
)

/* this function will change the value of
the variable named "fifty" in the main
function directly, after passing in
the variable "seven" which was assigned
to the address of the variable "fifty".
*/
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
