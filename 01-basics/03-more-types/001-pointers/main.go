package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701 // create 2 int values (`i` and `j`)

	p := &i         // `p` points to `i`'s memory address
	fmt.Println(*p) // read `i` through the pointer - 42
	*p = 21         // set `i` through pointer. // "*T" means "the value at", so in this case we are setting the value at `i`.
	fmt.Println(i)  // see updated value of `i` - 21

	// NOTE: The "&" operator generates a pointer to its operand.
	p = &j         // set `p` to point to `j` // since p was already of type "pointer to int" we do not violate any strong typing rules
	*p = *p / 37   //divide `j` through the pointer
	fmt.Println(j) // see the new value of `j`
}
