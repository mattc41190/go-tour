package main

import "fmt"

// Lesson: An untyped constant takes the type needed by its context.
// An int can store at maximum a 64-bit integer, and sometimes less.

const (
	Big   = 1 << 100  // 1 followed by 100 zeroes (binary)
	Small = Big >> 99 // 1 followed by 1 zero (binary) in decimal the value is 2
)

func needInt(i int) int { return i*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	// needInt(Big) // panics
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
