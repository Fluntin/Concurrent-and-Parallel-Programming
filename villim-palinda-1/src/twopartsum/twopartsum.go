package main

import (
	"fmt"
)

// sum the numbers in a and send the result on res.
// a is slice of integers
// res chan<- int, channel we use to send the sum
func sum(a []int, res chan<- int) {
	// TODO sum a
	sum := 0

	//index, value := range collection
	for _, element := range a {
		sum = sum + element
	}
	// TODO send result on res
	res <- sum
}

// concurrently sum the array a.
// splitting the array in half -> start two Goroutines to sum each half concurrently using the sum function.
func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)

	//concurrently sum -> sum1 then sum2
	//order in which the subtotals are received by the main Goroutine is not guaranteed ???

	//Slice that includes the first n/2 elements
	go sum(a[:n/2], ch)
	////Slice that includes the last n/2 elements
	go sum(a[n/2:], ch)

	// TODO Get the subtotals from the channel and return their sum
	//both Goroutines use the same channel
	subtotal1 := <-ch
	subtotal2 := <-ch
	return subtotal1 + subtotal2
}

func main() {
	// example call
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ConcurrentSum(a))
}
