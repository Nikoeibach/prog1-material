package listprops

import "fmt"

func ExampleCountValue() {
	list := []int{1, 3, 1, 4, 1, 5, 2, 5, 3, 5}
	emptyList := []int{}

	fmt.Println(CountValue(list, 1))
	fmt.Println(CountValue(list, 5))
	fmt.Println(CountValue(list, 4))
	fmt.Println(CountValue(list, 3))

	fmt.Println()

	fmt.Println(CountValue(emptyList, 1))
	fmt.Println(CountValue(emptyList, 5))
	fmt.Println(CountValue(emptyList, 4))
	fmt.Println(CountValue(emptyList, 3))

	// Output:
	// 3
	// 3
	// 1
	// 2
	//
	// 0
	// 0
	// 0
	// 0
}
