package listprops

import "fmt"

func ExamplePrimeValues() {
	list := []int{1, 2, 3, 4, 5}
	emptyList := []int{}

	fmt.Println(PrimeValues(list))
	fmt.Println(PrimeValues(emptyList))

	// Output:
	// 3
	// 0
}
