package tables

// Lookup erwartet zwei Listen und einen Wert.
// Sucht den Wert in der längeren Liste(l1) und liefert den entsprechenden Wert aus der anderen Liste(l2).
// Liefert einen leeren String, falls der gesuchte Wert
// in l1 nicht vorkommt oder falls die gefundene Position nicht in l2 vorkommt.
func Lookup(l1, l2 []string, v string) string {
	if len(l1) == 0 || len(l2) == 0 {
		return ""
	}

	if len(l1) > len(l2) {
		return LookupV(l1, l2, v)
	} else {
		return LookupV(l2, l1, v)
	}

}

func LookupV(liste1, liste2 []string, v string) string {

	for pos, el := range liste1 {
		if el == v {
			if pos < len(liste2) {
				return liste2[pos]
			} else {
				return ""
			}
		}

	}

	return ""

}

// for i := 0; i < len(l1); i++ {
// 	if l1[i] == v {

// 		if i < len(l2) {
// 			return l2[i]
// 		} else {
// 			return " "
// 		}

// 	}
// }

// return " "

// Lookup erwartet zwei Listen und einen Wert.
// Sucht den Wert in der längeren Liste(l1) und liefert den entsprechenden Wert aus der anderen Liste(l2).
// Liefert einen leeren String, falls der gesuchte Wert
// in l1 nicht vorkommt oder falls die gefundene Position nicht in l2 vorkommt.
// func LLookup(l1, l2 []string, v string) string {
// 	if len(l1) == 0 || len(l2) == 0 {
// 		return ""
// 	}
// 	for i := 0; i < len(l1); i++ {
// 		if l1[i] == v {
// 			return l2[i]
// 		}
// 	}

// 	return ""

// }
