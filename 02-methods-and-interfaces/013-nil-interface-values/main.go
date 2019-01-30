package main

import "fmt"

type I interface {
	M()
}

func main() {
	fmt.Println("CODE WILL PANIC -- READ CODE TO SEE WHY")
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
