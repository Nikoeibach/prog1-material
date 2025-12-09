package listprops

// CountValue gibt erwartet eine Liste von Zahlen und eine Zahl v.
// Liefert die Anzahl der Vorkommen von v in der Liste.
func CountValue(list []int, v int) int {
	count := 0
	listlength := len(list)

	for i := 0; i < listlength; i++ {
		if list[i] == v {

			count++
		}

	}

	return count
}
