package main

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/game"
	"github.com/RodrigoCelso/gophercises-10/player"
)

func main() {
	dealer := player.New()
	game := game.Game{Shoe: *deck.New(deck.WithMultipleDeckSize(4), deck.WithShuffle())}
	game.Hit(dealer)
	fmt.Println(dealer)
}
