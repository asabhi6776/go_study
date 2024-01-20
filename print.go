package main

import (
	"fmt"
)

func main() {
	// var city string = "Lucknow"
	// fmt.Println(city)

	// println for printing a new line

	var city string = "Lucknow"
	var user string = "Abhishek"

	var grades int = 42

	fmt.Print("Welcome to ", city, ", ", user, "\n")

	fmt.Print(city, "\n")
	fmt.Print(user, "\n")

	fmt.Printf("Nice to see you here, at %v", city)
	fmt.Printf("Marks: %d", grades)

}
