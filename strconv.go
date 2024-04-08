package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Converting integers to string
	var i int = 42
	var s string = strconv.Itoa(i)

	// Converting string to integer
	var a string = "200"
	c, err := strconv.Atoi(a)

	// Expecting errors
	var b string = "100abhishek"
	d, err := strconv.Atoi(b)

	fmt.Printf("%v, %T \n", c, c)
	fmt.Printf("%v, %T \n", err, err)
	fmt.Printf("%v, %T \n", d, d)
	fmt.Printf("%v, %T \n", err, err)
	fmt.Printf("%q", s)
}
