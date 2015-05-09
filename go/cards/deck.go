package deck

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
