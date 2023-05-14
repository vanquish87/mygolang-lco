package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	jasmeet := User{"Jasmeet", "myemail@your.server", true, 32}
	fmt.Println(jasmeet)
	fmt.Printf("jasmeet details are : %+v\n",jasmeet)
	fmt.Printf("Name is %v and email is %v",jasmeet.Name, jasmeet.Email)
	
}

// I think structs are sort of Class in python
// no inheritance in golang, no super or parent
type User struct {
	Name	string
	Email 	string
	Status	bool
	Age		int
}