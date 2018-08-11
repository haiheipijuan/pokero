package lib

// represents a single playing card.
type Card int

// the rank of the card
func (c Card) Rank() Rank {
	return Rank(c % 13)
}

// the suit of the card.
func (c Card) Suit() Suit {
	return Suit(c / 13)
}

func (c Card) String() string {
	return c.Rank().String() + c.Suit().String()
}