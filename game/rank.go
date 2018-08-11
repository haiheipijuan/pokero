package game

// represents the rank of a card.
type Rank int

const (
	Two   Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var ranksStr = "23456789TJQKA"

func (r Rank) String() string {
	return ranksStr[r : r+1]
}
