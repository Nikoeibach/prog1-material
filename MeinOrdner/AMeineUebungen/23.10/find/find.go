package tables

// Find erwartet eine Liste und einen Wert.
// Sucht den Wert in der Liste und liefert die Position.
// Liefert -1, falls der Wert nicht in der Liste vorkommt.

func Find(list []string, v string) int {
	if len(list) == 0 {
		return -1
	} else {
		for pos, el := range list {
			if el == v {
				return pos
			}

		}
	}
	return -1
}
