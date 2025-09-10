package game

import (
	"strconv"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/deck"
)

type Game struct {
	Shoe           deck.Deck
	DiscardTray    deck.Deck
	Dealer         *bjackplayer.Player
	Players        []*bjackplayer.Player
	PlayerMaxScore int
}

func New(playerQuantity int) *Game {
	var players []*bjackplayer.Player
	for idx := range playerQuantity {
		players = append(players, bjackplayer.New("Player"+strconv.Itoa(idx+1)))
	}
	return &Game{Shoe: *deck.New(deck.WithMultipleDeckSize(4), deck.WithShuffle()), Dealer: bjackplayer.New("Dealer"), Players: players}
}
