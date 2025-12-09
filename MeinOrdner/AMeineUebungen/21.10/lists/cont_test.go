package listprops

import "fmt"

func ExampleContains() {
	l := []string{"Hallo", "Welt", "gute", "Nacht"}
	fmt.Println(Contains(l, "Welt"))
	fmt.Println(Contains(l, "Haus"))

	// Output:
	// true
	// false
}

func ExampleContainsOnly() {
	l := []string{"Welt", "Welt", "Welt", "Welt"}
	x := "Welt"
	fmt.Println(ContainsOnly(l, x))

	// Output:
	// true
}

func ExampleContainsPos() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt"}
	x := "Welt"
	fmt.Println(ContainsPos(l1, x))

	//Output:
	// [1]
}

func ExampleContainsN() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt"}
	fmt.Println(ContainsN(l1, "Welt", 2))
	fmt.Println(ContainsN(l1, "Haus", 2))

	// Output:
	// true
	// false
}

func ExampleContainsNRow() {
	l1 := []string{"Welt", "Haus", "Welt", "Welt", "Welt"}
	fmt.Println(ContainsNRow(l1, "Welt", 2))
	fmt.Println(ContainsNRow(l1, "Welt", 3))
	fmt.Println(ContainsNRow(l1, "Welt", 4))

	// Output:
	// true
	// true
	// false
}

func ExampleContainsNMax() {
	// Liste: "Welt" kommt einmal einzeln vor, und einmal als 3er-Kette.
	// Die maximale Kettenlänge ist also 3.
	l1 := []string{"Welt", "Haus", "Welt", "Welt", "Welt", "Haus"}

	// Frage: Ist die längste Kette genau 2 lang? -> Nein (sie ist 3)
	fmt.Println(ContainsNMax(l1, "Welt", 2))

	// Frage: Ist die längste Kette genau 3 lang? -> Ja!
	fmt.Println(ContainsNMax(l1, "Welt", 3))

	// Frage: Ist die längste Kette genau 4 lang? -> Nein (nur 3)
	fmt.Println(ContainsNMax(l1, "Welt", 4))

	// Output:
	// false
	// true
	// false
}
