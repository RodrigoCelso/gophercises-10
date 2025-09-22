package game

import (
	"strconv"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/deck"
)

type Game struct {
	Shoe            deck.Deck
	DiscardTray     deck.Deck
	Dealer          *bjackplayer.Player
	Users           []*bjackplayer.Player
	NPCs            []*bjackplayer.Player
	NPCIntelligence float64
	PlayerMaxScore  int
}

func New(npcsQuantity int, users []*bjackplayer.Player) *Game {
	var npcs []*bjackplayer.Player
	for idx := range npcsQuantity {
		npcs = append(npcs, bjackplayer.New("Player_"+strconv.Itoa(idx)))
	}

	return &Game{
		Shoe:            *deck.New(deck.WithMultipleDeckSize(4), deck.WithShuffle()),
		Dealer:          bjackplayer.New("Dealer"),
		Users:           users,
		NPCs:            npcs,
		NPCIntelligence: 1,
	}
}
