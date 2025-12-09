package listprops

// ContainsN liefert true, falls die Liste l
// den String x mindestens n mal enthält.
func ContainsN(l []string, x string, n int) bool {
	var Count int
	for _, el := range l {
		if el == x {
			Count++
		}
	}
	if Count >= n {
		return true
	}
	return false
}

// Count := 0
// 	for i := 0; i < len(l); i++ {
// 		if l[i] == x {
// 			Count++
// 		}

// 	}
// 	if Count >= n {
// 		return true
// 	}
// 	return false
