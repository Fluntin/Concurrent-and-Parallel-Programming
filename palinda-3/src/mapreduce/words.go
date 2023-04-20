// TotalSeconds      : 2.6473839
package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freq := make(map[string]int)
	words := strings.Fields(text)
	// Added by me --------------------------------------
	subRoutines := 10 // same as julia I will make subRoutines and devide the work.
	wg := new(sync.WaitGroup)
	size := len(words)

	//receive the local frequency maps from each goroutine and merge them into a single frequency map in the main function
	frequency_chan := make(chan map[string]int)

	//Devide the file into several pieces and distribute them among the subRoutines.
	for i := 0; i < size; i += size / subRoutines {
		j := i + len(words)/subRoutines

		//Fix -> Error's for sum of pieces > file...
		if j > size {
			j = size
		}

		piece := words[i:j]
		wg.Add(1)
		// Now I have my pieces and WaitGroups i can start my subRoutines

		go func() {
			defer wg.Done()
			//This is a frequency map
			subResult := make(map[string]int)
			for _, word := range piece {
				word = strings.ToLower(word)
				word = strings.TrimFunc(word, func(r rune) bool {
					return !((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'))
				})
				if len(word) > 0 {
					subResult[word]++
				}

			}
			frequency_chan <- subResult
		}()
	}

	//Need to close the frequency_chan
	go func() {
		wg.Wait()
		close(frequency_chan)
	}()

	for {
		localFreq, ok := <-frequency_chan
		if !ok {
			break
		}
		for word, count := range localFreq {
			freq[word] += count
		}
	}

	// ---------------------------------------------------
	return freq
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data
	data, err := os.ReadFile(DataFile)
	if err != nil {
		panic(err)
	}
	//This is how i test but there has to be a way to use the test_words.go file
	fmt.Printf("%#v", WordCount(string(data)))
	//In Go, the "%#v" verb in the fmt.Printf() function is used to print the value
	//of a variable in a Go-syntax-like format. It is useful for debugging purposes,
	//as it shows the exact representation of a variable in Go syntax.

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)

}
