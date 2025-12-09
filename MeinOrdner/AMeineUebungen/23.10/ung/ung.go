package listprops

// OddValues gibt erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen in der Liste.
func OddValues(list []int) int {
	Count := 0
	for _, el := range list {
		if el%2 != 0 {
			Count++
		}
	}
	return Count
}
