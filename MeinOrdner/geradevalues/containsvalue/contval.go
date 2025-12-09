package listprops

// ContainsValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert true, falls v in der Liste enthalten ist, sonst false.
func ContainsValue(list []int, v int) bool {
	if len(list) == 0 {
		return false
	}

	for _, zahl := range list {
		if zahl == v {
			return true
		}
	}

	return false
}

// EvenValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der geraden Zahlen in der Liste.
func EvenValues(list []int) int {
	Count := 0

	for _, zahl := range list {
		if zahl%2 == 0 {
			Count++
		}
	}

	return Count
}

// OddValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen in der Liste.
func OddValues(list []int) int {
	Count := 0

	for _, zahl := range list {
		if zahl%2 != 0 {
			Count++
		}
	}

	return Count
}

// PrimeValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der Primzahlen in der Liste.
func PrimeValues(list []int) int {
	Count := 0
	if len(list) == 0 {
		return 0
	}

	for _, zahl := range list {
		if IsPrime(zahl) {
			Count++
		}
	}

	return Count
}

func IsPrime(n int) bool {
	teiler := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			teiler++
		}
	}

	return teiler == 2

}

//optimiert
// func IsPrime(n int) bool {
// if n<=1{
//return false}
// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}

// 	return true

// }
