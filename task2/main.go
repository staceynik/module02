package main

import (
	"fmt"
	"math"
)

func main() {
	velocityInMetersPerSecond := 120.4

	// Convert velocity from m/s to km/h
	velocityInKilometersPerHour := velocityInMetersPerSecond * 3.6
	europeanVelocity := math.Round(velocityInKilometersPerHour*100) / 100

	// Convert velocity from m/s to mph
	velocityInMilesPerHour := velocityInMetersPerSecond * 2.23694
	americanVelocity := math.Round(velocityInMilesPerHour*100) / 100

	fmt.Printf("European velocity: %.2f km/h\n", europeanVelocity)
	fmt.Printf("American velocity: %.2f mph\n", americanVelocity)
}
