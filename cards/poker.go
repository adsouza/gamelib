package cards

import "fmt"

type Category uint8

const (
	HighCard Category = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

const (
	FaceDown Status = iota
	FaceUp
)

type PokerState struct {
	faceDown, faceUp uint
}

func (ps *PokerState) NumCards(s Status) (uint, error) {
	switch s {
	case FaceDown:
		return ps.faceDown, nil
	case FaceUp:
		return ps.faceUp, nil
	default:
		return 0, fmt.Errorf("unknown status %d for poker", s)
	}
}

func (ps *PokerState) Add(s State) (State, error) {
	fd, err := s.NumCards(FaceDown)
	if err != nil {
		return ps, fmt.Errorf("cannot change state of card in hand: %s", err)
	}
	ps.faceDown += fd
	fu, err := s.NumCards(FaceUp)
	if err != nil {
		return ps, fmt.Errorf("cannot change state of card in hand: %s", err)
	}
	ps.faceUp += fu
	return ps, nil
}

var comparators map[Category]func(Hand, Hand) (int, error)

func (h Hand) Category() Category {
	return HighCard
}

func Compare(a, b Hand) (int, error) {
	if aCat, bCat := a.Category(), b.Category(); aCat != bCat {
		return int(bCat) - int(aCat), nil
	}
	return comparators[a.Category()](a, b)
}

func init() {
	comparators[HighCard] = compareHighCards
}

func compareHighCards(a, b Hand) (int, error) {
	aRanks, err := extractRanks(a)
	if err != nil {
		return 0, err
	}
	bRanks, err := extractRanks(b)
	if err != nil {
		return 0, err
	}
	for i, r := range aRanks {
		if r == bRanks[i] {
			continue
		}
		return int(bRanks[i]) - int(r), nil
	}
	return 0, nil
}

func extractRanks(h Hand) ([]Rank, error) {
	result := []Rank{}
	for r := King + 1; r >= Two; r-- {
		n, err := countCardsOfRank(h, r)
		if err != nil {
			return nil, err
		}
		if n > 1 {
			return nil, fmt.Errorf("hand has more than 1 card of the same rank")
		}
		if n == 1 {
			result = append(result, r)
		}
	}
	return result, nil
}

func countCardsOfRank(h Hand, r Rank) (uint, error) {
	n := uint(0)
	for _, s := range []Suit{Spades, Hearts, Clubs, Diamonds} {
		sn, err := h[Card{Rank: r, Suit: s}].NumCards(FaceUp)
		if err != nil {
			return 0, err
		}
		n += sn
	}
	return n, nil
}
