package main

import(
	"fmt"
)

func Sqrt(x float64) float64 {
	var prevGuess float64
	var guess = 1.0
	for i := 0; i < 10; i++ {
		if prevGuess == guess {
			return guess
		}
		prevGuess = guess
		guess -= (guess*guess - x) / (2*guess)
	}
	return guess
}

func main() {
	fmt.Println(Sqrt(4))
}

// First Loop: (Assuming 4.0 passed in)
// 1 x 1 - 4 = -3
// 2 * 1 = 2
// -3 / 2 = -1.5
// 1 - -1.5 = 2.5

// Second Loop:
// 2.5 * 2.5 - 4 = 2.25
// 2 * 2.5 = 5
// 2.25 / 5 = .45
// 2.5 - 0.45 = 2.05
// ...