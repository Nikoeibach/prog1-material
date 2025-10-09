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
		if NumberGood(guess, my_number) {
			fmt.Println("Das war richtig!")
			return
		}

		fmt.Printf("%v war nicht korrekt!\n", guess)
	}
	fmt.Println("Game Over!")
}

// Read Number liefert uns einen Int
func ReadNumber() int {
	var n int

	fmt.Print("Bitte gib eine Zahl ein: ")
	fmt.Scan(&n)

	return n
}

// NumberGood prüft ob x = einer zufällig gewählten zahl zwischen 1 und 100 ist
func NumberGood(x, n int) bool {

	//if x == my_number {
	//return true}
	//return false

	return x == n

}
