package dice

import (
	"math/rand"
	"time"
)

type Dice struct {
	sides, quantity uint8
	rng             rand.Rand
}

func New(sides, quantity uint8) Dice {
	return Dice{sides: sides, quantity: quantity, rng: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func NewStdPair() Dice {
	return New(6, 2)
}

func (d Dice) Roll() uint16 {
	var r uint16
	for i := 0; i < d.quantity; i++ {
		r += uint16(d.rng.Intn(d.sides) + 1)
	}
	return r
}
