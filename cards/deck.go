package cards

import (
	"fmt"
	"math/rand"
)

type Color int8
type Suit int8
type Rank int8
type Card struct {
	Suit Suit
	Rank Rank
}

func (c *Card) IsJoker() bool {
	return c.Suit == 0 && c.Rank == 0
}

// State is flexible and may be used to model the visibility of a card's face (i.e. up/down)
// or the number of identical cards in a multi-deck game.
type State int8
type Hand map[Card]State
type Deck struct {
	cards []Card
}

func (d *Deck) shuffle() {
	n := len(d.cards)
	if n < 2 {
		return
	}
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) Deal() (*Card, error) {
	if len(d.cards) == 0 {
		return nil, fmt.Errorf("no cards left in deck")
	}
	c := d.cards[0]
	d.cards = d.cards[1:]
	return &c, nil
}
