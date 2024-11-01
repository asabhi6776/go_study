package main
import "fmt"
func main() {
	arr := [5]int{5, 10, 15, 20, 25}
	i :=2
	slice_1 := arr[:i]
	slice_2 := arr[i+1:]
	slice := append(slice_1, slice_2...)
	fmt.Println(slice)
}