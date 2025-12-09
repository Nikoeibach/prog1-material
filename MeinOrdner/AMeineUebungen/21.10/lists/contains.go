package listprops

// Contains liefert true, falls die Liste l
// den String x enthält.
func Contains(l []string, x string) bool {
	for i := 1; i < len(l); i++ {
		if l[i] == x {
			return true
		}

	}
	return false
}
