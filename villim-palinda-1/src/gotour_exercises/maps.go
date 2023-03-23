// It should return a map of the counts of each “word” in the string s
// Count => How many time the word repeats in a given string...
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// return map -> keys are strings and values are integers.
func WordCount(s string) map[string]int {

	// I got strings.Fields() from a package "strings"
	// words is a slice o strings.
	words := strings.Fields(s)

	// Empty map to store the word counts
	counts := make(map[string]int)

	// Iterate over the words and increment their counts in the map
	// _, to ignore the index
	for _, word := range words {
		counts[word] += 1
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
