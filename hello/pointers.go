package main

import "fmt"

func verbose(){

	var number int = 42
	
	var address *int // declare an int pointer
	address = &number // assign memory address of number to address variable
	
	var value = *address // derefence value

	fmt.Printf("address %v\n", address) // memory address
	fmt.Printf("value: %v\n", value) // 42
}

func main(){
	verbose()
	
	number := 420 // int
	var address *int = &number // NOTE: we can't use `:=` syntax here. Why?
	value := *address

	fmt.Printf("address %v\n", address) // memory address
	fmt.Printf("value: %v\n", value) // 42
}