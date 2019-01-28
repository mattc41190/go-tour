package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prevNum := 0
	result := 0
	return func() int {
		if result == 0 {
			tmp := result
			result++
			return tmp
		} else {
			tmp := result
			result += prevNum
			prevNum = tmp
		}
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
