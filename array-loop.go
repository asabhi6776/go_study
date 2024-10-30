package main
import "fmt"

func main() {

	grades := [6]int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}
	
}