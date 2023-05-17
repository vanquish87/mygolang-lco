package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	// loginCount := 23
	// var loginCount = 23
	var loginCount int = 10

	var result string

	if loginCount < 10 {
		result = "regular user"
		} else if loginCount > 10 {
		result = "High user"
	} else {
		result = "Exactly 10 logins"
	}
 
	fmt.Println(loginCount)
	fmt.Println(result)

	if 9 % 2 == 0 {
		fmt.Println("Number is even")
	}	else {
		fmt.Println("Number is odd")
	}

	// web request handling simple example
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
		} else {
			fmt.Println("Num is NOT less than 10")
		}

	// error checking use like try except block
	// if err != nil {

	// }
}