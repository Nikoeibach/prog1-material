package searching

func BinFind(l []int, x int) int {
	links := 0
	rechts := len(l)
	for rechts >= links {

		mitte := (rechts-links)/2 + links // gibt die genaue Mitte auch bei ungerader Länge der Liste

		if x == l[mitte] {
			return mitte + links // müssen return nutzen weil man einen int ausgeben muss
		}

		if x < l[mitte] {

			rechts = mitte //nur den linken Teil übrig , alles von 0 bis exclusive mitte

		} else { //lasse nur den rechten Teil übrig
			// von mitte +1 bis zu len(l)
			links = mitte + 1
		}

	}
	return -1

}
