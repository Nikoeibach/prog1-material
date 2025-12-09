package listsearch

// Sucht x in l und liefert die Position des
// ersten Vorkommens von x, falls dies existiert.
// Falls x nicht in l vorkommt, wird -1 geliefert.
func Find(l []int, x int) int {
	for i := 1; i < len(l); i++ {
		if l[i] == x {
			return i
		}

	}
	return -1
}

// Sucht x in l und liefert eine Liste
// mit allen Vorkommen von x in l.
func FindAll(l []int, x int) []int {
	result := []int{}
	for i := 1; i < len(l); i++ {
		if l[i] == x {
			result = append(result, i)
		}

	}
	return result
}
