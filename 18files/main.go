package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("We are doing files in golang!")
	content := "Yes, almost everything is moved from master to slave."

	file, err := os.Create("./myfile.txt")

	checkNilErr(err)
	
	length, err := io.WriteString(file, content)

	writes := os.Wr
	
	checkNilErr(err)
	
	fmt.Println("Length is:", length)
	
	defer file.Close()
	
	readFile("./myfile.txt")
}

func readFile(filename string)  {
	databyte , err := ioutil.ReadFile(filename)
	
	checkNilErr(err)
	
	fmt.Println("Data bytes inside file are: \n", databyte)
	fmt.Println("Text inside file is: \n", string(databyte))
	
}

// to avoid repetition
func checkNilErr(err error)  {
	// reverse logic to get error, like try n catch block
	// panic shut downs the program
	if err != nil {
		panic(err)
	}
}