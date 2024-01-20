package main

import (
	"fmt"
)

func main() {
	var s, t string = "Abhishek", "Singh"
	fmt.Println(s)
	fmt.Println(t)

	var (
		a string = "Singh"
		b int    = 23
	)
	fmt.Println(a)
	fmt.Println(b)

	x := "Hello World!"
	x = "Hello World"
	fmt.Println(x)

	y := 12
	y = 11

	fmt.Println(y)

}
