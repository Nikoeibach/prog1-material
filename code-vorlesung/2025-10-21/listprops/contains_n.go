package listprops

// ContainsN liefert true, falls die Liste l
// den String x mindestens n mal enthält.
func ContainsN(l []string, x string, n int) bool {
	Count := 0

	if len(l) == 0 {
		return false
	}

	for _, el := range l {
		if el == x {
			Count++
		}
	}

	return Count >= n

}
