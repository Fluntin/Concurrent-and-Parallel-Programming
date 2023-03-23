package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	//first two numbers to start
	current, next := 0, 1

	// Closure function as return
	return func() int {

		result := current
		current, next = next, current+next

		return result
	}
}

func main() {
	// Call the fibonacci function to get a closure
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
