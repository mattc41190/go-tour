package main

import "fmt"

type I interface { // Interface -- A set of methods
	M()
}

type T struct { // Struct -- A composite type
	S string
}

func (t T) M() { // Method with receiver
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"} // The set of methods defined in I is met by T
	i.M()
}
