package main

import (
	"fmt"
)

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area/10.0
}

func main() {
	fmt.Printf("%0.2f liters needed.\n", paintNeeded(4.2, 3.0))
}
