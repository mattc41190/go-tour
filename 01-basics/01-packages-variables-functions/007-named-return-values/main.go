package main

import "fmt"

// split splits a 2 digit number into 2 values which when added together produce the original sum
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // <- naked return (can harm readability; use sparingly)
}

func main() {
	fmt.Println(split(17))
}
