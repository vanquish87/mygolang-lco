package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	
	// this is a pointer *http.Response
	fmt.Printf("Response is of type %T\n", response)
	
	// caller's responsibility to close the connection
	defer response.Body.Close()
	
	databyte, err := ioutil.ReadAll(response.Body)
	
	if err != nil {
		panic(err)
	}

	// get actual string data
	content := string(databyte)
	fmt.Println(content)
}