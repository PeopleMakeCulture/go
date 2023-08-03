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

}