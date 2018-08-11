package lib

// represents the suit of a card.
type Suit int

const (
	Diamonds Suit = iota
	Clubs
	Hearts
	Spades
)

var suitsStr = []rune{'\u2660', '\u2663', '\u2665', '\u2660'}

func (s Suit) String() string {
	return string(suitsStr[s])
}