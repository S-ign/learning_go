package main

import (
  "fmt"
  "time"
  "reflect"
)

func reportNap(name string, delay int) {
  for i:=0;i<delay;i++{
	fmt.Println(name, "sleeping")
	//causes main func to sleep
	time.Sleep(1*time.Second)
  }
  fmt.Println(name, "wakes up!")

}

func send(myChannel chan string) {
  reportNap("send goroutine", 2)
  fmt.Println("***sending value***")
  myChannel <- "a"
  fmt.Println("***sending value***")
  myChannel <- "b"
}

func main() {
  //create a channel
  myChannel := make(chan string)

  //start a goroutine(similar to thred)
  go send(myChannel)

  reportNap("main goroutine", 5)

  //get values sent to myChannel from send func
  fmt.Println(<-myChannel)
  fmt.Println(<-myChannel)

  fmt.Println(reflect.TypeOf(myChannel))
}

//***Output***

//main goroutine sleeping
//send goroutine sleeping
//main goroutine sleeping
//send goroutine sleeping
//send goroutine wakes up!
//***sending value***
//main goroutine sleeping
//main goroutine sleeping
//main goroutine sleeping
//main goroutine wakes up!
//a
//***sending value***
//b
