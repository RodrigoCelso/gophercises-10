package game

import "github.com/RodrigoCelso/gophercises-10/deck"

type Game struct {
	Shoe        []deck.Card
	DiscardTray []deck.Card
}
