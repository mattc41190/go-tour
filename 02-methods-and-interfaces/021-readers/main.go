package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello Reader!")

	b := make([]byte, 8) // make a byte slice with a length of 8 all sets to their 0 value. In this case, uint8(0). I think...

	for { // infinite loop
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) // Each time we go through the loop the byte slice is written over.
		if err == io.EOF {
			break
		}
	}
}
