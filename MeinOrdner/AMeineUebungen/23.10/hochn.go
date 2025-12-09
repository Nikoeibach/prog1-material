package power

// Pow2 berechnet "2 hoch n".
func Pow2(n int) int {

	if n == 0 {
		return 1
	}
	return 2 * Pow2(n-1)

}

func Hoch2(n int) int { //n ist der exponent
	if n == 0 {
		return 1
	}
	result := 1 // basisfall 2 hoch 0

	for i := 1; i <= n; i++ {
		result = result * 2
	}
	return result

}

// Power berechnet x hoch n (x^n) rekursiv.
// Beispiel: Power(5, 3) -> 5 * 5 * 5 = 125

func Power(x int, n int) int {
	if n == 0 {
		return 1
	}
	return x * Power(x, n-1)
}

// MaxList erwartet eine Liste von Zahlen und liefert die größte Zahl zurück.
// Du darfst davon ausgehen, dass die Liste niemals leer ist.
func MaxList(list []int) int {

	if len(list) == 1 {
		return list[0]
	}

	maxRest := MaxList(list[1:])

	if list[0] > maxRest {
		return list[0]
	} else {
		return maxRest
	}
}
