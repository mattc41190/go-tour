package main

import "fmt"

// Purpose of this lesson is to consider that when dealing with pointers to structs
// we do not have to do explicit dereferencing.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // p.X = 1e9 == (*p).X = 1e9 // <- point of confusion
	fmt.Println(v)
}
