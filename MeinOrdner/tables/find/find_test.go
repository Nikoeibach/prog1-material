package tables

import "fmt"

func ExampleFind() {
	l1 := []string{"Haus", "Apfel", "Holz", "Wald", "Auto", "Vorlesung", "Fahrrad"}
	l2 := []string{}

	fmt.Println(Find(l1, "Haus"))
	fmt.Println(Find(l1, "Apfel"))
	fmt.Println(Find(l1, "Vorlesung"))
	fmt.Println(Find(l1, "Birne"))

	fmt.Println()

	fmt.Println(Find(l2, "Haus"))
	fmt.Println(Find(l2, "Apfel"))
	fmt.Println(Find(l2, "Vorlesung"))
	fmt.Println(Find(l2, "Birne"))

	// Output:
	// 0
	// 1
	// 5
	// -1
	//
	// -1
	// -1
	// -1
	// -1
}
