package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Card struct {
	Suit  string
	Rank  string
	Value int
}

type Deck []Card

func (d Deck) Shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d *Deck) Deal() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

func NewDeck() Deck {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	deck := Deck{}
	for _, suit := range suits {
		for i, rank := range ranks {
			value := i + 1
			if value > 10 {
				value = 10
			}
			card := Card{Suit: suit, Rank: rank, Value: value}
			deck = append(deck, card)
		}
	}
	deck.Shuffle()
	return deck
}

func playerWins(playerCards []Card, dealerCards []Card) bool {
	playerTotal := 0
	for _, card := range playerCards {
		playerTotal += card.Value
	}
	dealerTotal := 0
	for _, card := range dealerCards {
		dealerTotal += card.Value
	}

	if playerTotal > 21 {
		return false
	} else if dealerTotal > 21 {
		return true
	} else {
		return playerTotal > dealerTotal
	}
}

func simulateGames() {
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Generate a unique seed value for the random number generator
			rand.Seed(time.Now().Unix() + int64(counter))
			counter++

			// Shuffle the deck and deal the cards
			deck := NewDeck()
			deck.Shuffle()
			playerCards := []Card{}
			playerCards = append(playerCards, deck.Deal())
			playerCards = append(playerCards, deck.Deal())
			dealerCards := []Card{}
			dealerCards = append(dealerCards, deck.Deal())
			dealerCards = append(dealerCards, deck.Deal())
			fmt.Println(playerCards)

			// Play the game
			for calculateHandValue(playerCards) < 17 {
				playerCards = append(playerCards, deck.Deal())
			}
			for calculateHandValue(dealerCards) < 17 {
				dealerCards = append(dealerCards, deck.Deal())
			}

			// Determine the winner
			if calculateHandValue(playerCards) > 21 {
				fmt.Println("Player busts")
			} else if calculateHandValue(dealerCards) > 21 {
				fmt.Println("Dealer busts")
			} else if calculateHandValue(playerCards) > calculateHandValue(dealerCards) {
				fmt.Println("Player wins")
			} else if calculateHandValue(playerCards) <= calculateHandValue(dealerCards) {
				fmt.Println("Dealer wins")
			} else {
				fmt.Println("Push")
			}
		}()
	}
	wg.Wait()
}
func calculateHandValue(hand []Card) int {
	numAces := 0
	value := 0
	for _, card := range hand {
		if card.Rank == "Ace" {
			numAces++
		} else {
			value += card.Value
		}
	}
	for i := 0; i < numAces; i++ {
		if value+11 <= 21 {
			value += 11
		} else {
			value++
		}
	}
	return value
}

func main() {
	/*deck := NewDeck()
	playerCards := []Card{}
	dealerCards := []Card{}
	playerCards = append(playerCards, deck.Deal())
	dealerCards = append(dealerCards, deck.Deal())
	playerCards = append(playerCards, deck.Deal())
	dealerCards = append(dealerCards, deck.Deal())
	*/
	simulateGames()
	// TODO: Implement the rest of the game mechanics.
}
