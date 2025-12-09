package listprops

// Contains liefert true, falls die Liste l
// den String x enthält.
func Contains(l []string, x string) bool {
	if len(l) == 0 {
		return false
	}
	for _, el := range l {
		if el == x {
			return true
		}
	}

	return false

}
