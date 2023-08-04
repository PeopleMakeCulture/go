package main

import "fmt"

/* Version 1: Integer Stack */

// define a stack
type intStack struct {
	idx int32
	elements [5]int32
}

// methods are like functions that take an addtnl argument before the function name
func (s *intStack) push (val int32) {
	s.elements[s.idx] = val
	s.idx ++
}

func (s *intStack) pop() int32 {
	s.idx--
	return s.elements[s.idx]
}

/* Version 2: Stack using Generics */

// NOTE: since we're using a generic we need to specify type of generic
type stack[T any] struct {
	elements [5]*T
	idx int32
}

func (t *stack[T]) t_push (val T) {
	var pointer *T = &val
	t.elements[t.idx] = pointer 
	t.idx++
} 

func (t *stack[T]) t_pop() *T {
	t.idx --
	return t.elements[t.idx]
}


func main(){
	s := new(intStack)
	s.push(6)
	s.push(8)
	last := s.pop()
	fmt.Printf("last: %v\n", last)
	fmt.Printf("stack: %v\n", *s) // {1 [123 4 0 0 0]} returns pointer, vals, and empty placeholders

	t := new(intStack)
	t.push(123)
	t.push(4)
	t_last := t.pop()
	fmt.Printf("last: %v\n", t_last)
	fmt.Printf("stack: %v\n", *t)
}
