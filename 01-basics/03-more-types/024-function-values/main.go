package main

import (
	"fmt"
	"math"
)

// compute accepts a function and returns the results of that functions invocation
// I know this is confusing... whomp
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))      // The function maybe called directly
	fmt.Println(compute(hypot))    // The function may be passed into another function for invocation
	fmt.Println(compute(math.Pow)) // Any func whomaches is allowed to be used. Generic programming... possibly ;)
}
