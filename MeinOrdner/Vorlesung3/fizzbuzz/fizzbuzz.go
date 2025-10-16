package fizzbuzz

import "fmt"

//FIZZ gibt die Zahlen von 1 bis 30 aus...
func FIZZ() {

	for i := 1; i <= 30; i++ {
		fmt.Println(i)

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")

		} else if i%3 == 0 {
			fmt.Println("fizz")

		} else if i%5 == 0 {
			fmt.Println("buzz")

		} else {
			fmt.Println(i)

		}

	}

}
