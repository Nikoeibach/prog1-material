package power

import "fmt"

func ExampleHoch2() {
	// Testfall 1: 2 hoch 0 (mathematisch: 1)
	fmt.Println(Hoch2(0))

	// Testfall 2: 2 hoch 1 (mathematisch: 2)
	fmt.Println(Hoch2(1))

	// Testfall 3: 2 hoch 5 (mathematisch: 32)
	fmt.Println(Hoch2(5))

	// Testfall 4: 2 hoch 10 (mathematisch: 1024)
	fmt.Println(Hoch2(10))

	// Output:
	// 1
	// 2
	// 32
	// 1024
}

func ExamplePow2() {
	// Testfall 1: 2 hoch 0 (Basisfall)
	fmt.Println(Pow2(0))

	// Testfall 2: 2 hoch 1
	fmt.Println(Pow2(1))

	// Testfall 3: 2 hoch 5
	fmt.Println(Pow2(5))

	// Testfall 4: 2 hoch 10
	fmt.Println(Pow2(10))

	// Output:
	// 1
	// 2
	// 32
	// 1024
}

func ExamplePower() {
	// 5 hoch 0 (Basisfall)
	fmt.Println(Power(5, 0))

	// 2 hoch 3 (2 * 2 * 2)
	fmt.Println(Power(2, 3))

	// 10 hoch 2 (10 * 10)
	fmt.Println(Power(10, 2))

	// Output:
	// 1
	// 8
	// 100
}

func ExampleMaxList() {
	// Test 1: Liste mit einem Element (Basisfall)
	fmt.Println(MaxList([]int{42}))

	// Test 2: Maximum in der Mitte
	fmt.Println(MaxList([]int{1, 5, 2, 3}))

	// Test 3: Maximum ganz am Anfang
	fmt.Println(MaxList([]int{10, 3, 1}))

	// Test 4: Maximum am Ende
	fmt.Println(MaxList([]int{2, 4, 8}))

	// Test 5: Negative Zahlen (Wichtig! Maximum ist -1, nicht 0)
	fmt.Println(MaxList([]int{-5, -10, -1}))

	// Output:
	// 42
	// 5
	// 10
	// 8
	// -1
}
