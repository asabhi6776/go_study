package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:2]

	arr_2 := [5]int{6, 7, 8, 9, 10}
	slice_2 := arr_2[:2]

	new_slice := append(slice, slice_2...)

	fmt.Println(new_slice)

}
