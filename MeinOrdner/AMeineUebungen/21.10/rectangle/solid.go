package rectangles

import "fmt"

// Erwartet zwei Seitenlängen `height` und `width`.
// Zeichnet ein Rechteck mit diesen Seitenlängen auf der Konsole.
// Das Rechteck soll komplett mit `#`-Zeichen gefüllt sein.
func DrawSolidRectangle(height, width int) {
	DrawRectangle(height, width, "#", "#")
}

func DrawRectangle(height, width int, inner, outer string) {

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if row == 0 || col == 0 || row == height-1 || col == width-1 {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}

		}
		fmt.Println()
	}

}

func DrawRightTriangle(size int, inner, outer string) { // es gibt auch rows und columns , die gleich size sind, mit jeder reihe , steigt die anzahl an columns +1
	for row := 0; row < size; row++ {
		for col := 0; col <= row; col++ {
			if col == 0 || col == row || row == size-1 {
				fmt.Print(outer)
			} else {
				fmt.Print(inner)
			}
		}
		fmt.Println()

	}
}
