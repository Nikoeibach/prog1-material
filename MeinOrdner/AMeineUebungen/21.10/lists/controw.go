package listprops

// ContainsNRow liefert true, falls die Liste l
// den String x n mal hintereinander enthält.
func ContainsNRow(l []string, x string, n int) bool { //l1, "Welt", 2
	counter := 0

	for _, el := range l {
		if el == x {
			counter++
		} else {
			if counter == n {
				return true
			}
			counter = 0
		}

	}
	return counter >= n //if counter == n {return true

}

func ContainsNMax(l []string, x string, n int) bool { //wie lang ist die längste kette von denselben strings
	maxcount := 0
	currentcount := 0

	for _, el := range l {
		if el == x {
			currentcount++
			if currentcount > maxcount {
				maxcount = currentcount
			}

		} else {
			currentcount = 0
		}
	}
	return maxcount == n

}
