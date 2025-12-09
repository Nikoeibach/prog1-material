package reservations

// ContainsBoth erwartet eine Liste von Strings und zwei Strings s1 und s2.
// Liefert true, falls sowohl s1 als auch s2 in der Liste enthalten sind, sonst false.
func ContainsBoth(list []string, s1 string, s2 string) bool {
	found1 := false
	found2 := false

	if len(list) == 0 {
		return false
	}

	for _, wort := range list {
		if wort == s1 {
			found1 = true
		}
		if wort == s2 {
			found2 = true
		}

	}

	return found1 && found2

}
