package main

import (
	"fmt"
)

// Lesson is heavily commented because I am not as smart as you

func sum(s []int, c chan int) {
	sum := 0              // create sum var
	for _, v := range s { // loop through s slice
		sum += v // add current value to the sum var
	}
	c <- sum // send the sum to the channel `c` // <- data flows in the direction of the arrow
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c) // pass second half of `s` and `c` to a `sum` goroutine
	go sum(s[len(s)/2:], c) // pass the first half of `s` and `c` to another `sum` goroutine

	// The first time data is pulled of c it goes to x (-9+4+0 == -5)
	// The second time data is pulled of c it goes to y (7+2+8 == 17)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
