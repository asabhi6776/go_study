// package main

// import "fmt"

// func main() {
// 	var grades [5]int
// 	fmt.Println(grades)
// 	var fruits [3]string
// 	fmt.Println(fruits)
// }

package main

import "fmt"

func main() {
	var grades [3]int = [3]int{10, 20, 30}
	grades2 := [4]int{5, 10, 20, 15}
	names := [...]string{"Abhishek", "Singh"}

	fmt.Println(grades)
	fmt.Println(grades2)
	fmt.Println(names)
}
