package main

import (
	"fmt"
	"math"
)

func main() {
	var length float64 = 35
	radius := length / (2 * math.Pi)
	area := math.Pi * math.Pow(*&radius, 2)
	area = math.Round(area*100) / 100 // rounding to two decimal places
	fmt.Printf("Radius of the circle: %.2f\n", radius)
	fmt.Printf("Area of the circle: %.2f\n", area)
}
