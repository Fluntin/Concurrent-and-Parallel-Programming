// Look at [`time.Format`](https://golang.org/pkg/time/#Time.Format)
package main

import (
	"fmt"
	"time"
)

// delay is of type time.Duration => scalar * time.Second
func Remind(text string, delay time.Duration) {
	for {
		// First I get the current time
		now := time.Now()

		//%02d is a format specifier for a decimal integer (d), with a minimum width of 2 digits
		//The 0 flag => any missing digits should be filled with zeros

		fmt.Printf("Current time is: %02d:%02d:%02d: %s\n", now.Hour(), now.Minute(), now.Second(), text)
		time.Sleep(delay)
	}
}

func main() {

	// execute the function in a separate Goroutine and allow the program to continue executing concurrently.
	go Remind("Time to eat", 10*time.Second)
	go Remind("Time to work", 30*time.Second)
	// This is main tread?
	Remind("Time to sleep", 60*time.Second)
	//To prevent the main program from exiting early, the following statement can be used: select{}
	select {}
}

//When a function is called as a statement, it is executed as a standalone statement and its return value (if any) is discarded.
//When a function is called as an expression, its return value is used as part of an expression.
