package library

// TYP-DEKLARATIONEN
type PageCount int // Seitenanzahl
type PriceCent int // Buchpreis

// KORREKTUR 1: Struct-Felder müssen mit Großbuchstaben beginnen,
// damit sie im Test verwendet werden können (Export).
type Book struct {
	Titel        string
	Autor        string
	Seitenanzahl PageCount
	Preis        PriceCent
}

// 1. Methode: IsLong() - Prüft, ob das Buch lang ist (> 500 Seiten).
func (b Book) IsLong() bool {
	// KORREKTUR 2: Nur das IF-Statement ist nötig, der Switch ist falsch.
	// Prüft das korrekte Feld 'b.Seitenanzahl'.
	if b.Seitenanzahl > 500 {
		return true // Wenn die Bedingung erfüllt ist, sofort true
	}

	return false // Ansonsten (nachdem die if-Bedingung nicht erfüllt wurde) false
}

// 2. Methode: GetPriceEuro() - Berechnet den Preis in Euro.
func (b Book) GetPriceEuro() float64 {
	// KORREKTUR 3: Berechnung und Rückgabe des float64-Werts.
	// Zuerst konvertieren wir den Preis (int) in float64, um eine Gleitkommadivision
	// (keine Ganzzahldivision) zu erzwingen.
	return float64(b.Preis) / 100.0
}
