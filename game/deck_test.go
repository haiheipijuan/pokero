package game

import (
	"testing"
)

// testing new deck and shuffle card cards in deck
func TestNewDeck(t *testing.T) {
	d := NewDeck()
	t.Logf("deck is created with the cards: %v", d.String())

	first := d.Cards[0]
	t.Logf("the first card in deck is: %v", first.String())

	d.ShuffleCards()
	t.Logf("cards in deck after shuffle: %v", d.String())

	sFirst := d.Cards[0]
	t.Logf("the first card in deck after shuffle is: %v", sFirst.String())

	if first == sFirst {
		t.Errorf("expected different card after shuffle but got same: %v <-> %v", first.String(), sFirst.String())
	}
}

// testing to pop a card in deck
func TestDeck_Pop(t *testing.T) {
	d := NewDeck()
	d.ShuffleCards()

	t.Logf("new cards in deck for testing pop method: %v", d.String())

	last := d.Cards[len(d.Cards) - 1]
	poppedCard := d.Pop()

	t.Logf("new cards in deck after pop: %v", d.String())

	if last != poppedCard {
		t.Errorf("expected %v but got %v", last, poppedCard)
	}

	nLast := d.Cards[len(d.Cards) - 1]
	if last == nLast {
		t.Errorf("mustnt be same card. %v != %v", last, nLast)
	}
}