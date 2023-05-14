package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Mango"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))
	
	var vegList = [3]string{"potato", "beans", "tomatoes"}
	fmt.Println(vegList)
	fmt.Println(len(vegList))
}