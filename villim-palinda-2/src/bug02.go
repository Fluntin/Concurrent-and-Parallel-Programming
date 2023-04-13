package main

import (
	"fmt"
	"sync"
	"time"
)

// This program should go to 11, but it seemingly only prints 1 to 10.
// We can do a waitgroup
func main() {
	//Unbuffered channel
	ch := make(chan int)
	var wait_group sync.WaitGroup
	wait_group.Add(1)

	//Error:"cannot use wait_group (variable of type sync.WaitGroup) as *sync.WaitGroup value in argument to Print"
	//I have to do pointer??? -> pointer points to a location in memory.
	go Print(ch, &wait_group)
	//1 - 11 in the channel
	for i := 1; i <= 11; i++ {
		ch <- i
	}

	//Print routine should complete -> Wait()
	// sohould wait before i close the ch but i get a Deadlock
	close(ch)
	wait_group.Wait() // sprijeci da se ostatk probgrama izvrsi dok counter ne postane 0!

}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
// Assumes that the channel will receive exactly 11 elements => 1 - 10
func Print(ch <-chan int, wait_group *sync.WaitGroup) {

	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	// Mark WaitGroup done when function returns
	defer wait_group.Done() // defer garatira da ce .Done() biti napravljen bez obzira na sve!

}

// Works !!!!
//Problem:
//For loop in the main function sends 11 to channel
//At this time Print function is still blocking on the previous element and will not get it since channel will close
//11 will die with ch when ch is closed ???.
