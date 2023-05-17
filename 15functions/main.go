package main

import "fmt"

func main() {
	fmt.Println("This is functions in Golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)

	fmt.Println("Sum is:", result)
	
	proRes, _ := proAdder(4, 6, 7, 8)
	fmt.Println("Pro Sum is:", proRes)
}

// doesn't matter how you declare sequence wise, even below also works
func greeterTwo() {
	fmt.Println("ANother method")
}
func greeter() {
	fmt.Println("Hello this is Golang bro!")
}

// return is int as function signature
func adder(valOne int, valTwo int) int {
	 return valOne + valTwo
} 

// if you need to return 2 things
func proAdder(values ...int) (int, string) {
	total := 0
	
	for _, val := range values {
		total += val
	}
	return total, "proAdder at work"
}