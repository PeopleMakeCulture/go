package main

import (
	"fmt"
	"reflect"
)

func main(){

	// static arrays are of fixed size
	languages := [9]string{
		"C", "Lisp", "C++", 
		"Java", "Python","Javascript", "Ruby", 
		"Go", "Rust", // trailing comma allows for multi-line list of values
	} 

	// slices can change size, but must refer to an underlying array
	classics := languages[:3] 
	new := languages[7:] 

	modern := make([]string, 4) // declare
	modern = languages[3:7] // assign

	allLangs := languages[:] // copy array to create a slice
	
	fmt.Printf("classics: %v\n", classics)
	fmt.Printf("new: %v\n", new)
	fmt.Printf("modern languages: %v\n", modern)
	fmt.Printf("type of languages: %v\n", reflect.TypeOf(languages).Kind())
	fmt.Printf("type of allLangs: %v\n", reflect.TypeOf(allLangs).Kind()) // slice

	// we can create a slice by not explicit stating size of array
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	fmt.Printf("type of frameworks: %v\n", reflect.TypeOf(frameworks).Kind()) // slice

	jsFrameworks := frameworks[0:4:4] // length 4 capacity 4

    // NOTE: append() creates a new slice of same type as original 
    // we need to re-assign the new array to the original variable
    frameworks = append(frameworks, "Meteor")

    fmt.Printf("js frameworks: %v\n", jsFrameworks)
    fmt.Printf("all frameworks updated: %v\n", frameworks)


}