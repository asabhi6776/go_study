package main

import (
	"fmt"
)

var job string = "Devops Engineer" // global variable

func main() {

	// Inner and outer block variable

	city := "Lucknow"
	{
		country := "India"
		fmt.Println(country)
		fmt.Println(city)
	}
	// fmt.Println(country)
	fmt.Println(city)

	name := "Abhishek" // Local variable

	fmt.Println(name)
	fmt.Println(job)

}
