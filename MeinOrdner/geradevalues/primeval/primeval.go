package listprops

// PrimeValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der Primzahlen in der Liste.
func PrimeValues(list []int) int {
	count := 0

	for i := 0; i < len(list); i++ {
		if IsPrime(list[i]) {
			count++
		}

	}
	return count
}

func IsPrime(n int) bool {

	if n <= 1 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
