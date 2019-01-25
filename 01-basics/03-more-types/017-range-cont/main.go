package main

import "fmt"

// Note: Think of the 1 on in the first loop as its binary representation i.e 1 == 00000001
// Then when we shift the value using bit shifting by `i`
// using `<<` 00000001 becomes that number of bits moved over.
// Example:
// 1 == 00000001
// 1 << uint(0) == 00000001 == 1
// 1 << uint(1) == 00000010 == 2
// 1 << uint(2) == 00000100 == 4
// 1 << uint(3) == 00001000 == 8
// 1 << uint(4) == 00010000 == 16
// 1 << uint(5) == 00100000 == 32
// 1 << uint(6) == 01000000 == 64
// 1 << uint(7) == 10000000 == 128

func main() {
	pow := make([]int, 10) // Create slice that refers to underlying array with 10 0s

	for i := range pow {
		pow[i] = 1 << uint(i) //
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
