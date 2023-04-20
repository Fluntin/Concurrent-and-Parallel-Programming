// UPPGIFT 3

// Denna är lite klurigare än förra uppgiften

// Här ska ni:
// 1. (Gjort) Skicka en sekvens av tal till ch1
// 2. (Gjort) Ta emot tal från ch1 och skicka kvadraten av talet till ch2
// 3. (Gjort) Ta emot tal från ch2 och skicka kuben av talet till ch3
// 4. Ta emot tal från ch2 och ch3 och skicka summan av talen till ch4
// 5. Ta emot tal från ch4 och skriv ut
// 6. Lägg till kod så att tal som tas emot av alla kanaler skrivs ut (alltså varje tal som tas emot från ch1, ch2 och ch3)
// 7. Ändra tidsintervallen och experimentera med hur det påverkar programmet

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch2toend := make(chan int)

	// Send a sequence of numbers to ch1
	go func() {
		for i := 1; i <= 10; i++ {
			ch1 <- i
			time.Sleep(10 * time.Millisecond)
		}
		close(ch1)
	}()

	// Square the numbers received from ch1 and send them to ch2
	go func() {
		for x := range ch1 {
			ch2 <- x * x
			ch2toend <- x * x
			time.Sleep(50 * time.Millisecond)
		}
		close(ch2)
	}()

	// Cube the numbers received from ch2 and send them to ch3
	go func() {
		numberruns := 0
		for x := range ch2 {
			ch3 <- x * x * x
			time.Sleep(100 * time.Millisecond)
			numberruns += 1

		}
		fmt.Println("done with Ch3", numberruns)
		close(ch3)
	}()

	// Combine the numbers received from ch2 and ch3 and send them to ch4
	go func() {
		for {
			select {
			case x, ok := <-ch2toend:
				if !ok {
					close(ch4)
					return
				}
				y := <-ch3
				ch4 <- x + y
			case y, ok := <-ch3:
				if !ok {
					close(ch4)
					return
				}
				x := <-ch2toend
				ch4 <- x + y
			}
			time.Sleep(75 * time.Millisecond)
		}
	}()
	for x := range ch4 {
		fmt.Println(x)
	}
	// Print the numbers received from ch4
	// TODO

}
