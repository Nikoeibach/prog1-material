package geradevalues

import "fmt"

func ExampleEvenValues() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	emptyList := []int{}

	fmt.Println(EvenValues(list))
	fmt.Println(EvenValues(emptyList))

	// Output:
	// 5
	// 0
}
