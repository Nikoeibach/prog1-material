package calc

// Liefert das Produkt von x und y.
func Mult(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	return x + Mult(x, y-1)
}

//5+Mult(5,4)
//5+Mult(5,3)
//5+Mult(5,2)
//5+Mult(5,1)
//5+Mult(5,0) Ende
