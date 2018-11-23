package cards

import (
	"math/rand"
	"testing"
	"time"
)

func TestDealing(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	d := NewDeck()
	n := len(d.cards)
	d.Deal()
	if len(d.cards) != n-1 {
		t.Fatalf("Dealing a card did not reduce the number of cards in the deck by 1.")
	}
}
