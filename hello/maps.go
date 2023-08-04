package main

import "fmt"

func main(){
	release_dates := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "Javascript": 1996, "Go": 2012,
	}

	for k, v := range release_dates {
		fmt.Printf("%s was first released in %d\n", k, v)
	}

	//accessing a non-existent key in a map will return the zero value of the value type
	fmt.Printf("Golang was first released in %d\n", release_dates["Golang"]) // 0

	// map[key] returns two values, the value, and a boolean for whether that key exists or not
	_ , exists := release_dates["Go"]

	fmt.Printf("It is %t that %s exists\n", exists, "Go")

	_ , exists = release_dates["Golang"]

	fmt.Printf("It is %t that %s exists\n", exists, "Golang")

}