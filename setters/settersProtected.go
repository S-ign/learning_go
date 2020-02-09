package main

import (
	"learning_go/setters/calandar"
	"fmt"
	"log"
)

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	date := calandar.Date{}
	err := date.SetYear(2020)

	//we're now unable to set:
	//date.Year = -50
	//because date.Year is unexported!

	logErr(err)

	err = date.SetMonth(2)
	logErr(err)

	err = date.SetDay(9)
	logErr(err)

	fmt.Println(date) //{2020, 2, 9}
}
