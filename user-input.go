package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey There, ", name)
	var nname string
	var is_muggle bool
	fmt.Print("Enter your Name & are you a muggle: ")
	fmt.Scanf("%s %t", &nname, &is_muggle)
	fmt.Println(name, is_muggle)

}
