package listsearch

import "fmt"

func ExampleFind() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23, 5}

	pos1 := Find(l1, 42)
	pos2 := Find(l1, 200)
	pos3 := Find(l1, 5)

	fmt.Println(pos1)
	fmt.Println(pos2)
	fmt.Println(pos3)

	// Output:
	// 2
	// -1
	// 1
}

func ExampleFindAll() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23, 5}

	pos1 := FindAll(l1, 42)
	pos2 := FindAll(l1, 200)
	pos3 := FindAll(l1, 5)

	fmt.Println(pos1)
	fmt.Println(pos2)
	fmt.Println(pos3)

	// Output:
	// [2]
	// [3]
	// [1 8]
}
