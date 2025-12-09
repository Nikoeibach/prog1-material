package searching

// / Liefert die Position von x in der Liste l,
// / oder liefert -1, falls x nicht in l vorkommt.
func Find(l []int, x int) int {
	for pos, el := range l {
		if el == x {
			return pos
		}
	}
	return -1
}
