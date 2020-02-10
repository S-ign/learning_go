package main

import (
	"fmt"
)

type Appliences interface {
	TurnOn()
}

type Fan string
func (f Fan) TurnOn(){
	fmt.Println(f)
}
func (f Fan) Max() {
	fmt.Println("Blowing everything away!")
}

type CoffeePot string
func (c CoffeePot) TurnOn(){
	fmt.Println(c)
}
func (c CoffeePot) Max() {
	fmt.Println(c)
}

func main() {
	var device Appliences = Fan("Power on!")
	device.TurnOn()

	//You can call methods of a type that are not a part of an
	//interface by surrounding the type with parenthesis.
	var fan Fan = device.(Fan)
	fan.Max()
}
