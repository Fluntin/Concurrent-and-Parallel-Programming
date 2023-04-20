package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := float64(1)           //  A decent starting guess for z is 1, no matter what the input.
	for i := 0; i < 10; i++ { // Repeat the calculation 10 times
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

// Next, change the loop condition to stop once the value has stopped changing
func Sqrt2(x float64) float64 {
	z_new := float64(1)
	for {
		z_old := z_new
		z_new -= (z_new*z_new - x) / (2 * z_new)
		if math.Abs(z_new-z_old) < 1e-10 {
			break
		}
	}
	return z_new
}

func main() {
	fmt.Println(Sqrt1(2))
	fmt.Println()
	fmt.Println(Sqrt2(2))
}
