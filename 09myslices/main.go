package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video in slices in Golang!")

	// []string is slice
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList %T\n", fruitList)
	fmt.Println(fruitList)
	
	// add to slice
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	
	highScores := make([]int, 4)
	
	highScores[0] = 123
	highScores[1] = 458
	highScores[2] = 253
	highScores[3] = 345
	// this will throw error because itni memory nai allocate kari thi 
	// highScores[4] = 555
	
	fmt.Println(highScores)

	// memory will be re-allocated
	highScores = append(highScores, 983, 543, 786, 611)
	fmt.Println(highScores)
	
	// sorting
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	
	// how to remove a value from slices based on index
	var courses = []string{"reactjs", "python", "java", "swift", "golang", "rust", "javascript"}
	fmt.Println(courses)
	
	var index int = 2
	// weird stuff
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}