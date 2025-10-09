package main

import (
	"fmt"
	"math/rand"
)

func main() {

	my_number := rand.Intn(100) + 1

	for i := 0; i < 3; i++ { //i++ heißt n=n+1
		guess := ReadNumber()

		// if guess == my_number  						//== compared 2 Variablen
		if guess < my_number {
			fmt.Println("Die gesuchte Zahl ist größer")
		}

		if guess > my_number {
			fmt.Println("Die gesuchte Zahl ist kleiner")

		}

		if guess == my_number {
			fmt.Println("Das war richtig!")
			return

		}

	}
	fmt.Println("Game Over! \nDie Zahl war:", my_number)

}

// Read Number liefert uns einen Int
func ReadNumber() int {
	var n int

	fmt.Print("Bitte gib eine Zahl ein: ")
	fmt.Scan(&n)

	return n
}
