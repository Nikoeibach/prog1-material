package listprops

import "fmt"

func ExampleOddValues() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	emptyList := []int{}

	fmt.Println(OddValues(list))
	fmt.Println(OddValues(emptyList))

	// Output:
	// 6
	// 0
}
