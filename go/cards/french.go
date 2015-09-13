package cards

const (
	_         = iota
	Red Color = iota
	Black
)
const (
	_          = iota
	♧ Suit = iota
	♢
	♡
	♤
)
const (
	_        = iota
	Ace Rank = iota
	Two
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
)

// DeckSize includes 2 jokers.
const DeckSize = 54

func NewDeck() Deck {
	d := Deck{}
	d.cards = make([]Card, DeckSize)
	c := 0
	for r := Ace; r <= King; r++ {
		for s := Clubs; s <= Spades; s++ {
			d.cards[c].Rank = r
			d.cards[c].Suit = s
			c++
		}
	}
	d.shuffle()
	return d
}
