package main

import "fmt"

// first letter Capital means a public usage type variable
const LoginToken string = "sadfetgsd"

func main() {
	var username string = "jasmeet"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	
	var smallFloat float32 = 255.65494948947984796549494894798479
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	
	var smallFloat1 float64 = 255.65494948947984796549494894798479
	fmt.Println(smallFloat1)
	fmt.Printf("Variable is of type: %T \n", smallFloat1)
	
	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)
	
	// nothing will print in garbage value of string
	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable is of type: %T \n", anotherStringVariable)
	
	// implicit type with the help of Lexer
	var website = "www.go.dev"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	
	// no var style -- only allowed inside methods not in Global
	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)
	
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
	
	

}