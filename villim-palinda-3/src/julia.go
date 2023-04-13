// Stefan Nilsson 2013-02-27
//Villim Prpic 2023-04-05

// Measure-Command { go run julia.go }
// Original-> TotalSeconds : 12.3817344
//
// Idea 1: separate each picture (0,1,2,3,...) itno its own go routine! -> TotalSeconds : 8.6561357
// Idea 2: separate each individual picture into small subpartitions and generate pixels between several goroutines.
// The number of separate go routines generating the image is refered to as magic_number.
//
//	with 10 subroutine generating pixels-> TotalSeconds      : 3.0076104
//	with 20 subroutine generating pixels-> TotalSeconds      : 2.795119
//	with 30 subroutine generating pixels-> TotalSeconds      : 2.87521
//	with 15 subroutine generating pixels-> TotalSeconds      : 2.8755093
//
// This program creates pictures of Julia sets (en.wikipedia.org/wiki/Julia_set).
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"strconv"
	"sync"
)

type ComplexFunc func(complex128) complex128

// takes a complex number and returns a complex number
// an array of functions...
// A programming language is said to have first-class functions if it treats functions as first-class citizens.
// This means the language supports:
// 1. passing functions as arguments to other functions
// 2. returning them as the values from other functions
// 3. assigning them to variables or storing them in data structures
// https://www.youtube.com/watch?v=kr0mpwqttM0&ab_channel=CoreySchafer

var Funcs []ComplexFunc = []ComplexFunc{
	//The Funcs slice contains a list of complex functions
	func(z complex128) complex128 { return z*z - 0.61803398875 },
	func(z complex128) complex128 { return z*z + complex(0, 1) },
	func(z complex128) complex128 { return z*z + complex(-0.835, -0.2321) },
	func(z complex128) complex128 { return z*z + complex(0.45, 0.1428) },
	func(z complex128) complex128 { return z*z*z + 0.400 },
	func(z complex128) complex128 { return cmplx.Exp(z*z*z) - 0.621 },
	func(z complex128) complex128 { return (z*z+z)/cmplx.Log(z) + complex(0.268, 0.060) },
	func(z complex128) complex128 { return cmplx.Sqrt(cmplx.Sinh(z*z)) + complex(0.065, 0.122) },
}

// Iterates through the Funcs slice and generates a PNG image for each function
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(len(Funcs))
	for n, fn := range Funcs {

		go CreatePng("picture-"+strconv.Itoa(n)+".png", fn, 1024, wg)

	}
	wg.Wait()
}

// CreatePng creates a PNG picture file with a Julia image of size n x n.
func CreatePng(filename string, f ComplexFunc, n int, wg *sync.WaitGroup) (err error) {
	defer wg.Done()
	file, err := os.Create(filename)
	//This pattern is often used in Go to handle errors in functions.
	if err != nil {
		return
	}

	defer file.Close()

	err = png.Encode(file, Julia(f, n))

	//copyed this from main()
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Julia returns an image of size n x n of the Julia set for f.
// We can separate this function!
func Julia(f ComplexFunc, n int) image.Image {
	wg := new(sync.WaitGroup) // This was me...
	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	img := image.NewRGBA(bounds)
	magic_number := 15
	for i := 0; i < magic_number; i++ {
		wg.Add(1)
		go divideImage(f, magic_number, i, n, img, wg)
	}
	wg.Wait()
	return img
}

func divideImage(f ComplexFunc, step, start, n int, img *image.RGBA, wg *sync.WaitGroup) {

	defer wg.Done()

	bounds := image.Rect(-n/2, -n/2, n/2, n/2)
	s := float64(n / 4)

	for i := bounds.Min.X + start; i < bounds.Max.X; i += step {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			n := Iterate(f, complex(float64(i)/s, float64(j)/s), 256)
			r := uint8(0)
			g := uint8(0)
			b := uint8(n % 32 * 8)
			img.Set(i, j, color.RGBA{r, g, b, 255})
		}
	}
}

// Iterate sets z_0 = z, and repeatedly computes z_n = f(z_{n-1}), n â‰¥ 1,
// until |z_n| > 2  or n = max and returns this n.
func Iterate(f ComplexFunc, z complex128, max int) (n int) {
	for ; n < max; n++ {
		if real(z)*real(z)+imag(z)*imag(z) > 4 {
			break
		}
		//The actual calculation whether to paint a pixel occurs in the `Iterate`
		//function in the _blink and you miss it_ line 73: `z = f(z)`,
		//using the function from the `Funcs` array that has been passed through several functions.
		z = f(z)
	}
	return
}
