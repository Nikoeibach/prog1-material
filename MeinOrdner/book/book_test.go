package library

import "fmt"

func ExampleBook() {
	// --- Testfall 1: Langes und teures Buch (über 500 Seiten) ---
	longBook := Book{
		Titel:        "Epic Go Adventures",
		Autor:        "Gopher Coder",
		Seitenanzahl: 650,  // Sollte IsLong() = true ergeben
		Preis:        2999, // Soll 29.99 Euro ergeben
	}

	// --- Testfall 2: Kurzes und günstiges Buch (unter 500 Seiten) ---
	shortBook := Book{
		Titel:        "Structs Quick Start",
		Autor:        "Niko",
		Seitenanzahl: 450, // Sollte IsLong() = false ergeben
		Preis:        995, // Soll 9.95 Euro ergeben
	}

	// Prüfe die IsLong() Methode
	fmt.Printf("Lang (Epic): %t\n", longBook.IsLong())
	fmt.Printf("Lang (Quick): %t\n", shortBook.IsLong())

	// Prüfe die GetPriceEuro() Methode (%.2f stellt sicher, dass zwei Nachkommastellen ausgegeben werden)
	fmt.Printf("Preis (Epic): %.2f €\n", longBook.GetPriceEuro())
	fmt.Printf("Preis (Quick): %.2f €\n", shortBook.GetPriceEuro())

	// Output:
	// Lang (Epic): true
	// Lang (Quick): false
	// Preis (Epic): 29.98 €
	// Preis (Quick): 9.95 €
}
