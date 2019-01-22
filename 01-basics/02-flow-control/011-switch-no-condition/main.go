package main

import (
	"fmt"
	"time"
)

// The expression `switch` with no condition is equivalent to `switch true {}`
// This means the first case which evaluates to true will be run.
// Drop in replamcement for long if-else chains.

func main() {
	t := time.Now()
	switch { // Notice the lack of expression // <- true
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
