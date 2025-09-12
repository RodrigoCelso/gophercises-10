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

func New(playerQuantity int, users []*bjackplayer.Player) *Game {
	var players []*bjackplayer.Player
	players = append(players, users...)

	for idx := range playerQuantity {
		players = append(players,
			&bjackplayer.Player{
				Name:     "Player" + strconv.Itoa(idx+1+len(users)),
				Playable: false,
			},
		)
	}

	return &Game{
		Shoe:    *deck.New(deck.WithMultipleDeckSize(4), deck.WithShuffle()),
		Dealer:  &bjackplayer.Player{Name: "Dealer", Playable: false},
		Players: players,
	}
}
