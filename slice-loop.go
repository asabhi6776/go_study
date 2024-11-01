package main
import "fmt"

func main(){

	arr := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println(index, "=>", value)
	}

	for _, value := range arr {
		fmt.Println(value)
	}

}