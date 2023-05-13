package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Burger :")

	// comma Ok // err err
	// since no try catch block in Go so error is caught as a Variable as 'err' or '_'
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}