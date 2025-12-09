package tables

import "fmt"

func ExampleLookup() {
	l1 := []string{"Haus", "Apfel", "Holz", "Wald", "Auto", "Vorlesung", "Fahrrad", "Birne"}
	l2 := []string{"house", "apple", "wood", "wood", "car", "lecture", "bicycle"}

	fmt.Println(Lookup(l1, l2, "Haus"))
	fmt.Println(Lookup(l1, l2, "Apfel"))
	fmt.Println(Lookup(l1, l2, "Vorlesung"))
	fmt.Println(Lookup(l1, l2, "Birne"))

	// Output:
	// house
	// apple
	// lecture
	//
}
