package main

import (
	"fmt"
	"gohelper"
	"strconv"
)

func main() {
	floats, err := gohelper.GetLines("floats")
	gohelper.HandleErr(err)
	sum := 0.0
	for _,sfloat := range floats {
		float, err := strconv.ParseFloat(sfloat, 64)
		gohelper.HandleErr(err)
		sum += float
	}

	fmt.Println(floats)
	fmt.Println(sum)
}
