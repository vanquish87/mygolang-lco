package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gasd4f4wrsz08"

func main() {
	fmt.Println("Welcome to handling urls in Golang!")
	fmt.Println(myurl)
	
	// parsing url
	result, _ := url.Parse(myurl)
	
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	
	qparams := result.Query()
	// url.Value is sort of saying its a dictionary
	fmt.Printf("Type of query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"][0])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	// constructing a url with info you have, in reverse order per se
	// need to pass reference (pointer to actual values) not the copies ie direct variables
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=jasmeet",
	}	
	
	// constructing URL using string method
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}