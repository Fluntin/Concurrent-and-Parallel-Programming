// För varje fråga skapa en ny fil och skriv koden i den filen.

// UPPGIFT 1

// Innan ni ändrar koden, skriv ner på papper vad problemet är, varför det uppstår och hur ni tror att det kan lösas.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for {
			select {
			case x, ok := <-ch1:
				if !ok {
					close(ch2)
					return
				}
				ch2 <- x * 2
			}
		}
	}()

	for x := range ch2 {
		fmt.Println(x)
	}
}
