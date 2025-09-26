package game

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/RodrigoCelso/gophercises-10/internal/deck"
)

type Game struct {
	Shoe               deck.Deck
	ShoeDecks          int
	DiscardTray        deck.Deck
	Dealer             *Player
	Users              []*Player
	NPCs               []*NPCPlayer
	NPCIntelligence    float32
	NPCTricksterChance float32
	CardCounter        int
	PlayerMaxScore     int
}

type GameOption func(Game) Game

func New(options ...GameOption) *Game {
	newGame := &Game{
		ShoeDecks:          1,
		Dealer:             NewPlayer("Dealer"),
		NPCIntelligence:    1,
		NPCTricksterChance: 0,
	}
	for _, option := range options {
		*newGame = option(*newGame)
	}
	if newGame.Shoe == nil {
		fmt.Println("shoe nil")
		newGame.Shoe = *deck.New(deck.WithShuffle())
	}
	return newGame
}

func WithShoeDecks(size int) GameOption {
	return func(g Game) Game {
		fmt.Println("shoe with decks")
		g.Shoe = *deck.New(deck.WithMultipleDeckSize(size), deck.WithShuffle())
		g.ShoeDecks = size
		return g
	}
}

func WithUsers(users []*Player) GameOption {
	return func(g Game) Game {
		g.Users = users
		return g
	}
}

func WithNPCs(npcQuantity int) GameOption {
	return func(g Game) Game {
		var npcs []*NPCPlayer
		for idx := range npcQuantity {
			isTrickster := rand.Float32() < g.NPCTricksterChance
			npcs = append(npcs, NewNPC("Player_"+strconv.Itoa(idx), isTrickster))
		}
		g.NPCs = npcs
		return g
	}
}

func WithNPCIntelligence(intelligence float32) GameOption {
	return func(g Game) Game {
		g.NPCIntelligence = intelligence
		return g
	}
}

func WithTrickster(chance float32) GameOption {
	return func(g Game) Game {
		if chance < 0 {
			chance = 0
		}
		g.NPCTricksterChance = chance

		// Define tricksters if npcs have been already created
		if g.NPCs != nil {
			for _, npc := range g.NPCs {
				npc.Trickster = rand.Float32() < g.NPCTricksterChance
			}
		}
		return g
	}
}
