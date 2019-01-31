package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e) // <-- causes infinite loop as %v will call the Error method recursively
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt find the square root of a positive number
func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var prevGuess float64
	var guess = 1.0
	for i := 0; i < 10; i++ {
		if prevGuess == guess {
			return guess, nil
		}
		prevGuess = guess
		guess -= (guess*guess - x) / (2 * guess)
	}
	return guess, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
