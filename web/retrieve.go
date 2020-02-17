package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// fetchURLBody returns the body of a URL and
// the size of the string in bytes.
func fetchURLBody(s string) (string, int) {
	response, err := http.Get(s)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body), len(body)
}


func main() {
	body, size := fetchURLBody("https://cnn.com")
	fmt.Println(body, size)
}
