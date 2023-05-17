package main

import "fmt"

func main() {
	// defer pushes is down to last to be executed in LIFO method just like Stack
	defer fmt.Println("This is last print")
	defer fmt.Println("World!") 
	defer fmt.Println("One") 
	defer fmt.Println("Two") 
	fmt.Println("Hello")
	myDefer()
}

// this will print in reverse as its Stack
func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}