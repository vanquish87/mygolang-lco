package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	
	presentTime := time.Now()
	fmt.Println(presentTime)
	
	// formatting date - very weird way
	fmt.Println(presentTime.Format("01-02-2006 PM Monday"))
	fmt.Println(presentTime.Format("2006-01-02 15:04:05 -07:00:00"))
	
	createdDate := time.Date(2020, time.March, 10, 14, 22, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 PM Monday"))
}