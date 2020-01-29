package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	//////Generate Rand number between 1-100
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	//////////////////////////////////////////

	//////Capture input from user

	//////Compare answer with target

	fmt.Println("I've chosen a random number between 1 and 100.\nCan you guess it?")
	success := false
	for x := 1; x <= 10; x++ {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		///////////////////////////

		//////Convert input to int
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		//////////////////////////

		if guess < target {
			fmt.Println("Oops. Your guess was LOW")
		}
		if guess > target {
			fmt.Println("Oops. Your guess was HIGH")
		}
		if guess == target {
			fmt.Println("You've got the correct answer!")
			success = true
			break
		}
	}
	if success == false {
		fmt.Println("Sorry you did not guess my number.")
	}
}
