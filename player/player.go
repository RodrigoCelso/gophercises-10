package player

import "github.com/RodrigoCelso/gophercises-10/deck"

type IPlayer interface {
	Hit()
	Stand()
	DoubleDown()
	Split()
}

type Player struct {
	Deck  []deck.Card
	Chips int
	Bet   int
}

func New() *Player {
	return &Player{}
}
