package main

import(
	"fmt"
)

func Sqrt(x float64) float64 {
	guess := 1.0
	for i := 0; i < 10; i++ {
		guess -= (guess*guess - x) / (2*guess)
		fmt.Println(guess)
	}
	return guess
}

func main() {
	Sqrt(4)
}