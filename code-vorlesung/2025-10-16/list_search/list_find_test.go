package listsearch

import "fmt"

// Sucht x in l und liefert die Position des
// ersten Vorkommens von x, falls dies existiert.
// Falls x nicht in l vorkommt, wird -1 geliefert.

func Find(l []int, x int) int {
	if len(l) == 0 {
		return -1
	}
	for pos, el := range l {
		if el == x {
			return pos
		}
	}
	return -1
}

// Sucht x in l und liefert eine Liste
// mit allen Vorkommen von x in l.
func FindAll(l []int, x int) []int {
	result := []int{}
	for pos, el := range l {
		if el == x {
			result = append(result, pos)

		}
	}
	return result
}

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
	// []
	// [1 8]
}
