package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Send a random number to one of the channels every 500ms
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			n := rand.Intn(3) + 1 // Randomly select a channel to send to
			switch n {
			case 1:
				ch1 <- n
			case 2:
				ch2 <- n
			case 3:
				ch3 <- n
			}
		}
	}()

	// Receive and print numbers from all channels using select
	count := make([]int, 3)

	for {
		// TODO Här ska ni skriva ett select statement som tar emot från alla kanaler
		//och räknar antalet tal som tas emot från varje kanal

		select {

		case value := <-ch1:
			fmt.Print(value)
			count[0] += 1

		case value := <-ch2:
			fmt.Print(value)
			count[1] += 1
		case value := <-ch3:
			fmt.Print(value)
			count[2] += 1

		}
		fmt.Print(count)

	}
}
