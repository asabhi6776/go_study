package main
import "fmt"

func main() {
	// slice := []{10, 20, 30}
	// fmt.Println(slice)
	// slice = append(slice, 40, 50)
	// fmt.Println(slice)

	arr := [4]int{10, 20, 30, 40}
	slice := arr[1:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = append(slice, 900, -90, 500)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}