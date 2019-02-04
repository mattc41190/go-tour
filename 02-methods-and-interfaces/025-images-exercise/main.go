package main

import (
	"fmt"
	"image"
	"image/color"
)

type MyImage struct {
}

// ColorModel returns the MyImage's color model.
func (img MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (img MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 2550, 2550)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (img MyImage) At(x, y int) color.Color {
	x8 := uint8(x)
	y8 := uint8(y)
	return color.RGBA{(x8 + y8) / 2, x8 * y8, x8 ^ y8, 255}
}

func main() {
	m := MyImage{}
	fmt.Println(m.At(2, 4))
	// pic.ShowImage(m) // <-- Only works in go tour
}
