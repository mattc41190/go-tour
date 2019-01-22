package main

import "fmt"

func main() {
	defer fmt.Println("world") // <- Run this line after all other statements in the function have been evaluated.

	fmt.Println("hello")
}
