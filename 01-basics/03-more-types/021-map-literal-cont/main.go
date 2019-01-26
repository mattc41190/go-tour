package main

import "fmt"

// Vertex represents a physical location in Lat & Lang
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},  // Note Vertex is not explicitly set. The Vertex type maybe assumed.
	"Google":    {37.42202, -122.08408}, // Note Vertex is not explicitly set. The Vertex type maybe assumed.
}

func main() {
	fmt.Println(m)
}
