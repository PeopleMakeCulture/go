package main

import (
	"fmt"
	"greetings"
)

func main() {
	fmt.Println("Hello World") // must use double quotes, single quotes are for one char rune literals
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
