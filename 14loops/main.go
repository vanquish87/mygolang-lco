package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Tuesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)
	
	// for d :=0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// more useful way
	// i gives you the index not Value
	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	// break, continue stuff
	rogueValue := 1

	for rogueValue < 12 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		// if rogueValue == 7 {
		// 	break
		// }
		if rogueValue == 9 {
			// goto a label
			goto lco
		}

		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}

	lco:
		fmt.Println("You are learning Golang")
}