package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://sowjsrkr.herokuapp.com/hello")
	if err != nil {
		fmt.Println(errors.New("error: request"))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(errors.New("error: read response"))
	}
	log.Println(string(body))
}
