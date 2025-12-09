package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// FilterDigits erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Ziffern entfernt werden.
// Alle anderen Zeichen sollen unverändert bleiben.
func FilterDigits(s string) string {
	result := "" // 1. Wir starten mit einem leeren String

	for _, el := range s {
		// 2. Wir prüfen: Ist das Zeichen 'el' KEINE Ziffer?
		// Eine Ziffer liegt zwischen '0' und '9'.
		// Also: Wenn es kleiner als '0' oder größer als '9' ist, behalten wir es.
		if el < '0' || el > '9' {

			// 3. Wir hängen das Zeichen an das Ergebnis an
			// Wichtig: string(el) wandelt den Rune zurück in Text um.
			result = result + string(el)
		}
	}

	return result
}
