package main

import "fmt"

func main() {

	var name string
	var age int
	var permission bool

	//alles einscannen +prüfen
	fmt.Print("Bitte geben sie ihren Namen ein: ")
	fmt.Scan(&name)

	fmt.Print("Bitte geben sie ihr Alter ein: ")
	fmt.Scan(&age)

	if age <= 10 {
		fmt.Printf("Hello %v you are just %v years old \nGET OUT!", name, age)
	} else if age > 10 && age < 18 {
		fmt.Print("Do you have permission from your parents (true/false)")
		fmt.Scan(&permission)

		if permission {
			fmt.Printf("Hello %v it is okay you can come in, even if you are just %v years old", name, age)
		} else {
			fmt.Printf("Hi %v, im sorry, but you have to get out", name)
		}
	} else {
		fmt.Printf("Welcome %v, you are %v years old, come in", name, age)
	}
}
