package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // send value of x to the channel
			x, y = y, x+y
		case <-quit: // quit can be pulled off the channel
			fmt.Println("quit")
			return // exit function
		}
	}
}

func main() {
	c := make(chan int)    // make a channel for ints
	quit := make(chan int) // make a channel for ints
	go func() {
		for i := 0; i < 10; i++ { // get first 10 numbers in the fib sequence
			fmt.Println(<-c) // pop off next value on the c channel
		}
		quit <- 0
	}() // create execute a go func
	fibonacci(c, quit) // runs before go func above finishes
}
