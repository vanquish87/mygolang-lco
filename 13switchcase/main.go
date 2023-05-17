package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in Golang")

	rand.Seed(time.Now().UnixNano())
	// generate number between 1 to 6
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is:", diceNumber)
	
	switch diceNumber {
	case 1:
		fmt.Println("You got 1 bro")
	case 2:
		fmt.Println("You got 2 bro")
	case 3:
		fmt.Println("You got 3 bro")
		fallthrough
	case 4:
		fmt.Println("You got 4 bro")
		fallthrough
	case 5:
		fmt.Println("You got 5 bro")
	case 6:
		fmt.Println("You got 6 bro")
	default:
		fmt.Println("What was that ??")
	}

}