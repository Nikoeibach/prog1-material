package fizzbuzz

import "fmt"

//FIZZ gibt die Zahlen von 1 bis 30 aus...
func FIZZImp(n, x, y int) {

	for i := 1; i <= n; i++ {

		if i%y == 0 && i%x == 0 {
			fmt.Println(i, "fizzbuzz")

		} else if i%x == 0 {
			fmt.Println(i, "fizz")

		} else if i%y == 0 {
			fmt.Println(i, "buzz")

		} else {
			fmt.Println(i)

		}

	}

}
