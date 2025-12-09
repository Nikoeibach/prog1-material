package reservations

// ContainsOverlap erwartet eine Liste von Strings und zwei Paare von jeweils zwei Strings.
// Liefert true, falls die String-Paare sich in der Liste überlappen,
// also falls z.B. t1 zwischen s1 und s2 liegt.
// Anmerkung: Die Reihenfolge der Strings im Paar ist hier nicht relevant.
// Anmerkung: An den Grenzen (also s1 == t1 oder s2 == t2) liegt kein Überlappen vor.
func ContainsOverlap(list []string, s1, s2, t1, t2 string) bool {
	poss1 := 0
	poss2 := 0
	post1 := 0
	post2 := 0

	if len(list) == 0 {
		return false
	}
	if Kommtvor(list, s1, s2, t1, t2) {
		for pos, el := range list {
			if el == s1 {
				poss1 = pos
			}
			if el == s1 {
				poss2 = pos
			}
			if el == s1 {
				post1 = pos
			}
			if el == s1 {
				post2 = pos
			}
		}

	}
	return post1 < poss2 && post1 > poss1 || post2 < poss2 && post2 > poss1

}

func Kommtvor(liste []string, s1, s2, t1, t2 string) bool {
	string1 := false
	string2 := false
	tstring1 := false
	tstring2 := false

	for _, el := range liste {
		if el == s1 {
			string1 = true
		}
		if el == s2 {
			string2 = true
		}
		if el == t1 {
			tstring1 = true
		}
		if el == t2 {
			tstring2 = true
		}

	}
	return string1 && string2 && tstring1 && tstring2
}
