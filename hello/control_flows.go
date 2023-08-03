package main

import "fmt"

// for loops
func for_loop(){
	// NOTE: go doesn't use parens around looping condition

	var DeploymentOptions = [4]string{"AWS", "GCP", "Azure", "R-pi"}

	fmt.Println("Option 1: Every other value")
	for i:= 0; i < len(DeploymentOptions); i+=2 {
		option := DeploymentOptions[i]
		fmt.Println(i, option)
	}

	fmt.Println("Option 2: For loop using range")
	for idx, val := range DeploymentOptions {
		fmt.Println(idx, val)
	}
}

// while loops are a subset of for loops
func while_loop(){
	count := 1
	for count < 5 {
		count += count
	}
	fmt.Printf("Final count: %d\n", count)
}

// function signatures contain a list of inputs and outputs
func find_average(x []float64)(avg float64){
	
	total := 0.0

	if len(x) == 0 {
		avg = 0
	} else {
		// since we don't use the index, we can't name it, or go will complain
		for _ , value := range x {
			total += value
		}

		avg = total / float64(len(x)) // we have to cast length as same type as avg
	}

	return // we don't have to say what to return here
}



func main(){

	for_loop()
	while_loop()

	// NOTE: Because find_avg() takes a slice as its argument, we can't make x a sized array here
	x := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Printf("Average val of x %v is %f", x, find_average(x))

}