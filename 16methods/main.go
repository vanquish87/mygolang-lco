package main

import "fmt"

func main() {
	fmt.Println("This is methods in Golang")
	jasmeet := User{"Jasmeet", "myemail@your.server", true, 32}
	fmt.Println(jasmeet)
	fmt.Printf("jasmeet details are : %+v\n",jasmeet)
	fmt.Printf("Name is %v and email is %v\n",jasmeet.Name, jasmeet.Email)
	
	jasmeet.GetStatus()
	// this only passes a copy and that's why we have pointers
	jasmeet.NewMail()
	// this will print same as above because only copy of data 'Email' was changed above
	fmt.Println(jasmeet)
	
}

// I think structs are sort of Class in python
// no inheritance in golang, no super or parent
type User struct {
	Name	string
	Email 	string
	Status	bool
	Age		int
}

// declaring a method by passing struct as argument in func
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

// change state of data of struct
func (u User) NewMail() {
	u.Email = "neweremail@magi.com"
	fmt.Println("Is user new email is:", u.Email)

}