package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // Has type Vertex
	v2 = Vertex{X: 1}  // T Vertex -- Y == 0
	v3 = Vertex{}      // X == 0 & Y ==0
	p  = &Vertex{1, 2} // Type "*Vertex"
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
