package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x        // send value of x on to the channel
		x, y = y, x+y // first time through x and y == 1
	}
	close(c)
}

func main() {
	c := make(chan int, 11) // Find eleventh number in fib
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
