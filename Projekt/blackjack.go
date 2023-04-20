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
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := range d {
		j := r.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
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
	var wg sync.WaitGroup
	wg.Add(len(suits))

	cardsChan := make(chan Card, len(suits)*len(ranks))

	for _, suit := range suits {
		go func(suit string) {
			defer wg.Done()
			for i, rank := range ranks {
				value := i + 1
				if value > 10 {
					value = 10
				}
				card := Card{Suit: suit, Rank: rank, Value: value}
				cardsChan <- card
			}
		}(suit)
	}

	wg.Wait()
	close(cardsChan)

	for card := range cardsChan {
		deck = append(deck, card)
	}

	//deck.Shuffle()
	return deck
}

type Player struct {
	Name    string
	Balance int
	Hand    []Card
}

func (p *Player) Hit(deck *Deck) {
	card := deck.Deal()
	p.Hand = append(p.Hand, card)
}

func (p *Player) Stand() {
	// Do nothing, the player is standing
}

func (p *Player) CalculateHandValue() int {
	value := 0
	hasAce := false
	for _, card := range p.Hand {
		value += card.Value
		if card.Rank == "Ace" {
			hasAce = true
		}
	}
	if hasAce && value+10 <= 21 {
		value += 10
	}
	return value
}

func NewPlayer(name string, balance int) Player {
	player := Player{Name: name, Balance: balance}
	player.Hand = make([]Card, 0)
	return player
}

type Table struct {
	Players []*Player
	Dealer  *Player
	Deck    Deck
}

func (t *Table) AddPlayer(player *Player) {
	t.Players = append(t.Players, player)
}

func (t *Table) Deal() {
	t.Deck = NewDeck()
	for i := 0; i < 2; i++ {
		for _, player := range t.Players {
			player.Hit(&t.Deck)
		}
		t.Dealer.Hit(&t.Deck)
	}
}

func (t *Table) Play() {
	for _, player := range t.Players {
		for player.CalculateHandValue() < 17 {
			player.Hit(&t.Deck)
		}
	}
	for t.Dealer.CalculateHandValue() < 17 {
		t.Dealer.Hit(&t.Deck)
	}
}

func NewTable() *Table {
	table := &Table{Players: make([]*Player, 0)}
	table.Dealer = &Player{Name: "Dealer", Balance: 0}
	return table
}

func simulateAlwaysHitPlayer(table *Table) bool {
	player := NewPlayer("Always Hit", 100)
	table.AddPlayer(&player)
	table.Deal()
	table.Play()
	playerScore := player.CalculateHandValue()
	dealerScore := table.Dealer.CalculateHandValue()
	if playerScore > 21 {
		return false
	} else if dealerScore > 21 || playerScore > dealerScore {
		return true
	} else {
		return false
	}

}
func runSimulations(numRuns int) int {
	tables := make([]*Table, 25)
	for i := 0; i < 25; i++ {
		table := NewTable()
		tables[i] = table
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	wins := 0

	for j := 0; j < numRuns; j++ {
		for _, table := range tables {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if simulateAlwaysHitPlayer(table) {
					mu.Lock()
					wins++
					mu.Unlock()
				}
			}()
		}
	}

	wg.Wait()
	return wins
}

func main() {
	deck := NewDeck()
	for _, card := range deck {
		fmt.Println(card)
	}
}
