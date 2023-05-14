package main

import "fmt"

func main() {
	fmt.Println("Maps in golang : Hash maps or Dict")
	
	// key is string as well as value is also string in map
	languages := make(map[string]string)
	
	languages["JS"] = "JavaScipt"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	
	fmt.Println(languages)
	fmt.Println("Value for JS is :", languages["JS"])

	// deleting
	delete(languages, "RB")
	fmt.Println(languages)

	// loops are interesting in GOlang
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// if you want to ignore the key just use '_'
	for _, value := range languages {
		fmt.Printf("value is %v\n", value)
	}
}