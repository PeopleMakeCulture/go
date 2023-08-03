package main

import "fmt"

// declare single var
var a int 

// declare group
var (
	b bool
	c float32
	d string
)

func main_old(){
	a = 42
	b, c = true, 32.0
	d = "string"
	fmt.Println(a, b, c, d)
}

func main(){
	a_new := 42
	b_new, c_new := true, 32.0 
	d_new := "string"
	fmt.Println(a_new, b_new, c_new, d_new)
}