package main

import "fmt"

func main() {

	src_slice := []int{1, 2, 3, 4, 5}
	dest_slice := make([]int, 3)

	num := copy(dest_slice, src_slice)

	fmt.Println(dest_slice)

	fmt.Println("Number of Element Copied:", num)

}
