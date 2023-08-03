package main

import "fmt"

func static_array() {
	// Static arrays are of fixed size
	// define a string array of size 4 that stores deployment options
	var DeploymentOptions = [4]string{"AWS", "GCP", "Azure", "R-pi"}

	/*FOR LOOPS*/
	// NOTE: go doesn't use parens around looping condition
	// (Option 1)
	fmt.Println("For loop the javascripty way")
	for i:= 0; i < len(DeploymentOptions); i++ {
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

	fmt.Printf("classics: %v\n", classics)
	fmt.Printf("new: %v\n", new)

	modern := make([]string, 4) // declare
	modern = languages[3:7] // assign

	fmt.Printf("modern languages: %v\n", modern)
}

func main(){
	static_array()
	slice()
}