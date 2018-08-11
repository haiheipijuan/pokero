package game

// represents the suit of a card.
type Suit int

const (
	Diamonds Suit = iota
	Clubs
	Hearts
	Spades
)

var suitsStr = []string{"♦", "♣", "♥", "♠"}

func (s Suit) String() string {
	return suitsStr[s]
}