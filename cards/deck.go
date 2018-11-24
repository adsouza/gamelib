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

// Status enumerates the various possible statuses that a given card may have (e.g. face-down, face-up, held).
type Status int8
type State interface {
	NumCards(Status) int
	Add(State) State
}
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

func (d *Deck) DealHands(handSize, numHands int, startState State) ([]Hand, error) {
	if startState == nil {
		return []Hand{}, fmt.Errorf("startState arg is nil.")
	}
	if handSize*numHands > len(d.cards) {
		return []Hand{}, fmt.Errorf("deck contains only %d cards so cannot deal %d cards to %d players.",
			len(d.cards), handSize, numHands)
	}
	hands := make([]Hand, numHands)
	for h := 0; h < numHands; h++ {
		hand := map[Card]State{}
		for c := 0; c < handSize; c++ {
			card := d.cards[c]
			hand[card] = startState.Add(hand[card])
		}
		hands[h] = hand
		d.cards = d.cards[handSize:]
	}
	return hands, nil
}
