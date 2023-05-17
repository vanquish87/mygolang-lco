package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("this is performing GET in golang! .. hitting a local Django API")

	myurl := "http://127.0.0.1:8000/api/tweets/"
	PerformGetRequest(myurl)
}

func PerformGetRequest(myurl string)  {
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	// close the connection : MANDATORY at the end via defer
	defer response.Body.Close()
	
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)
	
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
	
	// doing different way using Builder
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is:", byteCount)

	fmt.Println(responseString.String())

}