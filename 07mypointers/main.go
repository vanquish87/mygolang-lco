package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// creating a pointer - use '*'
	var ptr *int
	fmt.Println("Value of pointer is", ptr)
	
	number := 34
	// referencing a pointer use '&'
	var point = &number
	fmt.Println("Value of reference pointer is", point)     
	// see inside pointer
	fmt.Println("Value of reference pointer is", *point)     
	
	*point = *point * 2
	fmt.Println("Value of reference pointer is", *point)     

}