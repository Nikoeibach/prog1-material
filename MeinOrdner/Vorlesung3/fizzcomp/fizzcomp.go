package fizzbuzz

import "fmt"

//FIZZ gibt die Zahlen von 1 bis 30 aus...
func FIZZComp() {

	for i := 1; i <= 30; i++ {

		//wenn weder durch 5 noch durch 3 teilbar ist
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue

		}
		if i%3 == 0 {
			fmt.Println("fizz")

		}
		if i%5 == 0 {
			fmt.Println("buzz")

			fmt.Println()
		}

	}
}
