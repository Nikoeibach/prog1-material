package listprops

// ContainsOnly liefert true, falls die Liste l
// ausschließlich den String x enthält.
func ContainsOnly(l []string, x string) bool {
	for _, el := range l {
		if el != x {
			return false
		}
	}
	return true
}
