package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}
	fmt.Println(slice)

	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	slice_1 := arr[1:8]
	fmt.Println(slice_1)

	sub_slice := slice_1[0:4]
	fmt.Println(sub_slice)
}
