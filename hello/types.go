package main

import "fmt" // why do we need quotes around package names to import?

/*lets declare some more vars with types*/
func main(){

	// User specified types
	const a int32 = 12 // int32s are called runes in go
	const b float32 = 20.5
	var c complex128 = 1 + 4i
	var d uint16 = 14 // 16-bit unsigned int

	// Go also handles implicit typing!
	n := 42 // rune
	pi := 3.14 // 32-bit float
	x, y := true, false // bool - note, no caps
	z := "go rocks" // string

	// Here's how we do string interpolation!
	// NOTE: go will not let you compile code w/ unused variables 

	fmt.Printf("user-specified types: %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default types:%T %T %T %T %T\n", n, pi, x, y, z)
}