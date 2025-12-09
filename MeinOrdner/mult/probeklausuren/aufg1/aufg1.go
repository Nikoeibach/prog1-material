package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	result := []int{}
	if len(list) == 0 {
		return result
	}

	for _, el := range list {
		if el >= 10 {
			return result
		}
		if el < 10 {
			result = append(result, el)
		}

	}
	return result
}
