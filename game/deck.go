package game

import (
	"math/rand"
	"strings"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() *Deck {
	rand.Seed(time.Now().UnixNano())

	var cards []Card
	for i := 0; i < 52; i++ {
		cards = append(cards, Card(i))
	}

	return &Deck{
		Cards: cards,
	}
}

// shuffles the cards with Fisher-Yates algorithm
func (d *Deck) ShuffleCards() {
	for j := len(d.Cards) - 1; j > 0; j-- {
		k := rand.Intn(j + 1)
		d.Cards[j], d.Cards[k] = d.Cards[k], d.Cards[j]
	}
}

// pops the last card from deck
func (d *Deck) Pop() Card {
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}

// returns the cards in string format
func (d *Deck) String() string {
	var s []string
	for _, card := range d.Cards {
		s = append(s, card.String())
	}

	return strings.Join(s, ",")
}