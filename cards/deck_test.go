package cards

import (
	"fmt"
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

const (
	FaceUp Status = iota
	FaceDown
)

type TestState struct {
	faceUp, faceDown uint8
}

func (ts TestState) NumCards(s Status) (uint8, error) {
	switch s {
	case FaceUp:
		return ts.faceUp, nil
	case FaceDown:
		return ts.faceDown, nil
	default:
		return 0, fmt.Errorf("card status %v is not supported by TestState.", s)
	}
}

func (ts TestState) Add(s State) State {
	if s == nil {
		return ts
	}
	ts.faceUp, _ = s.NumCards(FaceUp)
	ts.faceDown, _ = s.NumCards(FaceDown)
	return ts
}

func TestDealingInitialHands(t *testing.T) {
	d := NewDeck()
	if _, err := d.DealHands(0, 0, nil); err == nil {
		t.Fatalf("Trying to use a nil interface for startState was not caught.")
	}
	hands, err := d.DealHands(6, 9, TestState{})
	if err != nil {
		t.Fatalf("Unable to deal out entire 54 card deck to 6 players.")
	}
	if len(hands) != 6 {
		t.Errorf("Wrong number of hands: got %d, want 6.", len(hands))
	}
}
