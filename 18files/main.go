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
	
	checkNilErr(err)
	
	fmt.Println("Length is:", length)
	
	defer file.Close()
	
	readFile("./myfile.txt")
}

func readFile(filename string)  {
	databyte , err := ioutil.ReadFile(filename)
	
	checkNilErr(err)

	// databyte is like :
	// [89 101 115 44 32 97 108 109 111 115 116 32 101 118 101 114 121 116 104 105 110 103 32 105 115 32 109 111 118 101 100 32 102 114 111 109 32 109 97 115 116 101 114 32 116 111 32 115 108 97 118 101 46]
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