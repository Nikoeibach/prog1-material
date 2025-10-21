package rectangles

import "fmt"

func DrawSolidRectangle(height, width int) {

	for row := 0; row < height; row++ {

		for col := 0; col < width; col++ {

			if row == 0 || col == 0 || row == height-1 || col == width-1 {
				fmt.Print("#")

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}
}
