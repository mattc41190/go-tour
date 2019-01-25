package main

import (
	"golang.org/x/tour/pic"
)

// Pic is draws pictures in golang tour
func Pic(dx, dy int) [][]uint8 {

	var result = [][]uint8{}

	for i := 0; i < dy; i++ {
		row := []uint8{}
		for j := 0; j < dx; j++ {
			row = append(row, uint8((j+i)*2))
		}
		result = append(result, row)
	}

	return result
}

func main() {
	// Must be run in
	pic.Show(Pic)
}
