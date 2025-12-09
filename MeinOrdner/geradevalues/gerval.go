package geradevalues

func EvenValues(list []int) int {
	count := 0

	for i := 0; i < len(list); i++ {
		if list[i]%2 == 0 {
			count++
		}

	}

	return count

}
