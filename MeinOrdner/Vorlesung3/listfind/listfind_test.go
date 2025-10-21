package listsearch

import "fmt"

// sucht x in l und liefert die Position.
// falls x nicht in l vorkommt wird -1 geliefert

// Die auskommentierte Funktion Find bleibt unverändert.

// func Find(l []int, x int) int {
//     for i := 0; i < len(l); i++ {
//         if l[i] == x {
//             return i
//         }
//     }
//     return -1
// }

// Korrigierte Funktion FindAll
func FindAll(l []int, x int) []int {
	result := []int{}
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			result = append(result, i)
		}

	}

	// Das return-Statement muss NACH der Schleife stehen, um die gesamte Liste zu durchsuchen.
	return result
}

func ExampleFindAll() {
	l1 := []int{17, 5, 42, 25, 3, -4, 8, -23, 5}

	pos1 := FindAll(l1, 42)
	pos2 := FindAll(l1, 5)  // Sucht nach 5, findet es bei Index 1 und 8
	pos3 := FindAll(l1, 99) // Erwartet [], daher Suche nach einer nicht vorhandenen Zahl (z.B. 99)

	fmt.Println(pos1)
	fmt.Println(pos2)
	fmt.Println(pos3)

	// Output:
	// [2]
	// [5]
	// [6]
}
