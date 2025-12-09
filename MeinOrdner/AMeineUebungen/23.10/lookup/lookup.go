package tables

// Lookup erwartet zwei Listen und einen Wert.
// Sucht den Wert in l1 und liefert den entsprechenden Wert aus l2.
// Liefert einen leeren String, falls der gesuchte Wert
// in l1 nicht vorkommt oder falls die gefundene Position nicht in l2 vorkommt.
func Lookup(l1, l2 []string, v string) string {
	pos := Find(l1, v)

	if pos == -1 || pos >= len(l2) {
		return ""
	} else {
		return l2[pos]
	}

}

func Find(list []string, v string) int {
	if len(list) == 0 {
		return -1
	}

	for pos, el := range list {
		if el == v {
			return pos
		}
	}
	return -1
}
