package main

import (
	"fmt"
)

func main() {
	var name string = "Abhishek Singh"
	var i int = 87
	fmt.Printf("Hey, %v! You have scored %d/100 in Chemistry", name, i)

	fmt.Printf("I am using %T datatype for grades", i)

	fmt.Printf("Printing my name %q in quote", name)
}
