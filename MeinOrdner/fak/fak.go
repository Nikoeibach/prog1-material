package main

import "fmt"

func main() {

	var n int

	fmt.Print("Gebe eine positive gerade Zahl ein!")
	fmt.Scan(&n)

	result := 1

	for i := 1; i <= n; i++ {
		result = result * i

	}

	fmt.Println(n, "!=", result)
}
