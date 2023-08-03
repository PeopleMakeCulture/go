package main

import (
	"fmt"
	"reflect"
)

func static_array() {
	// Static arrays are of fixed size
	// define a string array of size 4 that stores deployment options
	var DeploymentOptions = [4]string{"AWS", "GCP", "Azure", "R-pi"}

	/*FOR LOOPS*/
	// NOTE: go doesn't use parens around looping condition
	// (Option 1)
	fmt.Println("Every other value")
	for i:= 0; i < len(DeploymentOptions); i+=2 {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}

	// (Option 2)
	fmt.Println("For loop using range")
	for idx, val := range DeploymentOptions {
		fmt.Println(idx, val)
	}
}

/*
NOTE: slices must refer to an underlying array
*/
func slice(){
	// define a static string array of size 9
	// trailing comma allows for multi-line list of values 
	languages := [9]string{
		"C", "Lisp", "C++", 
		"Java", "Python","Javascript", "Ruby", 
		"Go", "Rust",
	} 

	classics := languages[:3] // [0:3]
	new := languages[7:] // [7:9]


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

		// go is not indent sensitive
		fmt.Printf("type of frameworks: %v\n", reflect.TypeOf(frameworks).Kind())

	jsFrameworks := frameworks[:4:4] //length 4 capacity 4
    
    fmt.Printf("js frameworks: %v\n", jsFrameworks)
    fmt.Printf("all frameworks: %v\n", frameworks)

    // NOTE: append() creates a new slice of same type as original 
    // we need to re-assign the new array to the original variable
    frameworks = append(frameworks, "Meteor")

    fmt.Printf("all frameworks updated: %v\n", frameworks)

}

func main(){
	static_array()
	slice()
}