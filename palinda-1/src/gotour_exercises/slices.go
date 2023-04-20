package main

import (
	"golang.org/x/tour/pic"
)

// dx and dy as input parameters => return 2-dimensional slice [][] of unit8
// uint8 is a built-in unsigned integer type in Go that occupies 8 bits of memory
// and can hold integer values between 0 and 255

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	// make takes type and length
	//In the case of a 2-dimensional slice, the length of the slice is the number of rows

	for y := range pic {

		//iterate over the elements of the outer slice pic and assign each element's index to the variable y

		pic[y] = make([]uint8, dx)
		for x := range pic[y] {

			// iterate over the elements of the inner slice pic[y] and assign each element's index to the variable x.

			pic[y][x] = uint8(x ^ y) // Nice!
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
