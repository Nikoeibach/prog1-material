package aufgabe3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountOdd.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein.
*/

// CountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
func CountOdd(list []int) int {
	Count := 0

	if len(list) == 0 {
		return 0
	}

	if list[0]%2 != 0 {
		Count++
	}

	return Count + CountOdd(list[1:])
}

//1+ CountOdd(list[1234])
//1+0+CountOdd([234])
//1+0+1+CountOdd([34])
//1+0+1+0+CountOdd([4])
//1+0+1+0+1+ CountOdd([])
//3+0

// func Ung(list []int) int {
// 	if len(list) == 0 {
// 		return 0
// 	}
// 	Count := 0
// 	for _, el := range list {
// 		if el%2 != 0 {
// 			Count++
// 		}
// 	}

// 	return Count
// }
