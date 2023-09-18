package main

import (
	"fmt"
	"os"
	"strings"
)

// classic for loop
func echo1(args []string){	
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

// for loop over range
func echo2(args []string){
	s, sep := "",""
	for _, arg := range args[1:] {
		s += sep + arg 
		sep = " "
	}
	fmt.Println(s)
}

// more efficient join
func echo3(args []string){
	fmt.Println(strings.Join(args[1:], " "))
}

// EXERCISES
// EX 1.1: Modify echo to print os.Args[0], the name of the command that invoked it
func ex1(args []string){
	fmt.Println(args[:1])
}
// EX 1.2: Modify echo to print the idx and val of each of its args, one per line
func ex2(args []string){
	for idx, arg := range args {
		fmt.Printf("idx: %d arg: %s\n", idx, arg)
	}
}
// TODO: EX 1.3: Experiment to measure diff in run time btwn inefficient versions and one that uses trings.Join (see 1.6 for `time` and 11.4 for benchmark tests)


func main() {
	echo1(os.Args)
	echo2(os.Args)
	echo3(os.Args)
	ex1(os.Args)
	ex2(os.Args)
}
