package game

import (
	"strconv"

	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/player"
)

type game struct {
	Shoe        deck.Deck
	DiscardTray deck.Deck
	Dealer      *player.Player
	Players     []*player.Player
}

func New(playerQuantity int) *game {
	var players []*player.Player
	for idx := range playerQuantity {
		players = append(players, player.New("Player"+strconv.Itoa(idx+1)))
	}
	return &game{Shoe: *deck.New(deck.WithMultipleDeckSize(4), deck.WithShuffle()), Dealer: player.New("Dealer"), Players: players}
}
