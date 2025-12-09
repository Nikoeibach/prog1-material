package listprops

// ContainsNRow liefert true, falls die Liste l
// den String x n mal hintereinander enthält.
func ContainsNRow(l []string, x string, n int) bool { //l1, "Welt", 2
	current := 0
	high := 0
	if len(l) == 0 {
		return false
	}

	for _, el := range l {
		if el == x {
			current++

			if current >= high {
				high = current
			}
		} else {
			current = 0
		}
	}

	return high == n
}
