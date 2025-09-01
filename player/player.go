package player

import (
	"math/bits"

	"github.com/RodrigoCelso/gophercises-10/deck"
)

type IPlayer interface {
	Hit()
	Stand()
	DoubleDown()
	Split()
}

type Player struct {
	Hand  []deck.Card
	Chips int
	Bet   int
}

func New() *Player {
	return &Player{}
}

func (p *Player) Score() int {
	var aces uint8
	var scoreValue int
	for _, card := range p.Hand {
		cardValue := bits.TrailingZeros16(card.Value) + 1
		switch cardValue {
		case 1:
			aces++
			scoreValue += 11
		case 11, 12, 13:
			scoreValue += 10
		default:
			scoreValue += cardValue
		}
	}
	for range aces {
		if scoreValue > 21 {
			scoreValue -= 10
			aces--
		} else {
			aces--
		}
	}
	return scoreValue
}
