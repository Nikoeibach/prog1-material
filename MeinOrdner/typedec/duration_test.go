package main

import "fmt"

func ExampleDuration_conversions() {
	fmt.Println(FromMinutes(60))
	fmt.Println(FromHours(3).Minutes())

	//Output:
	//3600

}
