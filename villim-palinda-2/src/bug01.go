package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)

	go func() {
		//Send the string "Hello world!" into the channel...
		ch <- "Hello world!"
	}()

	//Deadlock -> data has to be received by another goroutine..
	fmt.Println(<-ch)
}

//Channels are shared objects that let goroutines communicate
// Fix : Separate goroutine! -> Closure
// Note: Channel has to have a type.
// Go Routine has to send and Go Routine has to recive
