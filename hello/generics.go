package main

import "fmt"

// [T comparable] says the type T should be comparable
// (arr []T) says this function takes one argument: a slice of any type
func get_first_element[T comparable](sl []T) T {
	
	return sl[0]
}

func main() {

	// T = int32
    intSlice := []int32{5, 2, 9, 1, 7}
    minInt := get_first_element(intSlice)
    fmt.Println("first_int_element:", minInt)

    // T = float64
    floatSlice := []float64{3.14, 1.618, 2.718}
    minFloat := get_first_element(floatSlice)
    fmt.Println("first_float_element:", minFloat)

    // T = string
    stringSlice := []string{"apple", "banana", "orange"}
    minString := get_first_element(stringSlice)
    fmt.Println("first_string_element:", minString)
}