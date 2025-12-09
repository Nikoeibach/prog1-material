package rectangles

import "fmt"

func ExampleDrawSolidRectangle() {
	DrawSolidRectangle(3, 3)
	fmt.Println()
	DrawSolidRectangle(3, 6)
	// Output:
	// ###
	// ###
	// ###
	//
	// ######
	// ######
	// ######

}

func ExampleDrawRightTriangle() {
	// Wir testen ein Dreieck der Größe 6.
	// Rand: "#", Füllung: "+"
	DrawRightTriangle(6, "+", "#")

	// Output:
	// #
	// ##
	// #+#
	// #++#
	// #+++#
	// ######
}

func ExampleDrawRightTriangle_small() {
	// Ein kleinerer Test der Größe 3, um die Logik zu prüfen
	DrawRightTriangle(3, " ", "X")

	// Output:
	// X
	// XX
	// XXX
}
