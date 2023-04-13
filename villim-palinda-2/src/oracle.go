// Villim 2023-03-30

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// ---------------------------------------------------------------X
// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	// TODO: Answer questions.
	go answer(questions, answers)
	// TODO: Make prophecies.
	go prophecy(answers)
	// TODO: Print answers.
	go print(answers)
	return questions
}

// ---------------------------------------------------------X
// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
// Takes in 2 parameters: questions thats a string and answer thats a chanel
// but i have to fix it to take in two chanels...
// I have to move longestWord to a separate function...
func prophecy(answers chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)
	// Cook up some pointless nonsense.
	nonsense := []string{
		"Know thyself.",
		"Beware the Ides of March.",
		"The wolf shall dwell with the lamb.",
		"You will go to the land of the Persians.",
		"When the pig learns to fly, Athens will be great.",
		"The sons of Cronos shall rule all the land.",
	}
	for {
		// Take random nonsense and put it into the answers chan
		time.Sleep(time.Duration(7+rand.Intn(22)) * time.Second) // Take some time
		answers <- "Apollo says " + nonsense[rand.Intn(len(nonsense))]
	}
}

// For each question in questions chan generate a prophecy and send it to answers chan
func answer(questions <-chan string, answers chan<- string) {
	for question := range questions {
		go response(question, answers)
	}
}

func response(question string, answers chan<- string) {

	expected_words := map[string]string{
		"what":  "What or where do you dare?",
		"when":  "Now or never!",
		"when?": "Never!",
		"who":   "Your best firend!",
		"who?":  "Your best firend!",
		"why?":  "Because I can!",
		"why":   "I know but i wont say...",
	}
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, word := range words {
		if expected_words[strings.ToLower(word)] != "" {
			answers <- expected_words[strings.ToLower(word)]
			return
		}
	}
	// Find the longest word.
	longestWord := ""
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}
	answers <- "Let me think about " + longestWord
}
func print(answer <-chan string) {
	for question := range answer {
		fmt.Print("Pythia speeks: ")
		slow_answer := []rune(question)
		for index := range question {
			time.Sleep(time.Duration(3+rand.Intn(33)) * time.Millisecond)
			fmt.Print(string(slow_answer[index]))

		}
		fmt.Print("\n") // add a new line
	}
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
