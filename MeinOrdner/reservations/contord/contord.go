package reservations

// ContainsInOrder erwartet eine Liste von Strings und zwei Strings s1 und s2.
// Liefert true, falls sowohl s1 als auch s2 in der Liste enthalten sind, und s1 vor s2 kommt.
func ContainsInOrder(list []string, s1 string, s2 string) bool {
	found1 := false
	pos1 := 0

	found2 := false
	pos2 := 0

	if len(list) == 0 {
		return false
	}

	for pos, wort := range list {
		if wort == s1 {
			found1 = true
			pos1 = pos
		}
		if wort == s2 {
			found2 = true
			pos2 = pos
		}
	}

	return found1 && found2 && pos1 < pos2

}
