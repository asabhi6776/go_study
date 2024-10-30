package main

import "fmt"

func main() {

	var fruits [5]string = [5]string{"apple", "mango", "orange", "banana", "grapes"}
	fmt.Println(fruits[2])
	fruits[3] = "papaya"
	fmt.Println(fruits)

}
