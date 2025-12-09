package listprops

// ContainsOnly liefert true, falls die Liste l
// ausschließlich den String x enthält.
func ContainsOnly(l []string, x string) bool {
	for _, el := range l {
		if el != x {
			return false
		}
	}
	return true

}

func ContainsPos(l []string, x string) []int {
	result := []int{}
	counter := 0
	for pos, el := range l {
		if el == x {
			result = append(result, pos)
			counter++
		}
	}
	if counter >= 1 {
		return result
	}
	return []int{}

}

// for i := 0; i < len(l); i++ {
// 		if l[i] != x {
// 			return false
// 		}
// 	}
// 	return true
