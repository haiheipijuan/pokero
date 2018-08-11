package lib

import (
	"testing"
)

var testRanks = []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
var testSuits = []Suit{Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Diamonds, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Clubs, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Hearts, Spades, Spades, Spades, Spades, Spades, Spades, Spades, Spades, Spades, Spades, Spades, Spades, Spades}

// card values should follow the table below.
// +---+-----+-------+------+------+-----+-------+-------+------+-----+------+-------+------+-----+
// |   | Two | Three | Four | Five | Six | Seven | Eight | Nine | Ten | Jack | Queen | King | Ace |
// +---+-----+-------+------+------+-----+-------+-------+------+-----+------+-------+------+-----+
// | ♦ |  0  |   1   |   2  |   3  |  4  |   5   |   6   |   7  |  8  |   9  |   10  |  11  |  12 |
// +---+-----+-------+------+------+-----+-------+-------+------+-----+------+-------+------+-----+
// | ♣ |  13 |   14  |  15  |  16  |  17 |   18  |   19  |  20  |  21 |  22  |   23  |  24  |  25 |
// +---+-----+-------+------+------+-----+-------+-------+------+-----+------+-------+------+-----+
// | ♥ |  26 |   27  |  28  |  29  |  30 |   31  |   32  |  33  |  34 |  35  |   36  |  37  |  38 |
// +---+-----+-------+------+------+-----+-------+-------+------+-----+------+-------+------+-----+
// | ♠ |  39 |   40  |  41  |  42  |  43 |   44  |   45  |  46  |  47 |  48  |   49  |  50  |  51 |
// +---+-----+-------+------+------+-----+-------+-------+------+-----+------+-------+------+-----+
func TestCardsByValue(t *testing.T) {
	for i := 0; i < 52; i++ {
		card := Card(i)

		t.Logf("card is testing %v", card.String())

		if testRanks[i] != card.Rank() {
			t.Errorf("expected %v but got %v", testRanks[i], card.Rank())
		}

		if testSuits[i] != card.Suit()  {
			t.Errorf("expected %v but got %v", testSuits[i], card.Suit())
		}
	}
}
