package main

import (
	"fmt"
	"errors"
	"log"
)

//combines three defined underlined types to create date structure
type Date struct {
	Year    int
	Month   int
	Day     int
}

////////////DATE SETTERS////////////////////////////////////
//setter methods that validates arguments.
//marking the reciever (*Date) a pointer is important
//since the passed in argument will be lost since
//we are not returning it.
func (d *Date) SetYear(year int) error {
	if year < 2000 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}
////////////////////////////////////////////////////////////

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	date := Date{}
	err := date.SetYear(2020)
	//****************************************
	//although our setter method helps with validation
	//we can still directly alter the struct directly!
	date.Year = -50
	//we must put our Date type in its own package
	//as an unexported struct.
	//****************************************
	logErr(err)

	err = date.SetMonth(2)
	logErr(err)

	err = date.SetDay(9)
	logErr(err)

	fmt.Println(date) //{-50, 2, 9}
}
