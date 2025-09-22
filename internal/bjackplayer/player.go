package bjackplayer

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/internal/deck"
)

type Player struct {
	Name      string
	Chips     int
	MainHand  PlayerHand
	SplitHand PlayerHand
	Splitted  bool
}

type PlayerHand struct {
	Cards            deck.Deck
	Bet              int
	NaturalBlackjack bool
}

func New(name string) *Player {
	return &Player{Name: name}
}

func (p *Player) String() string {
	profile := fmt.Sprint(p.Name, ":\nHand cards\n")
	for _, card := range p.MainHand.Cards {
		profile += fmt.Sprintln("-", card)
	}
	if p.Splitted {
		profile += fmt.Sprintln("\nSecond hand cards")
		for _, card := range p.SplitHand.Cards {
			profile += fmt.Sprintln("-", card)
		}
	}
	profile += fmt.Sprint("Hand score:", p.MainHand.Cards.BlackjackScore())

	if p.Splitted {
		profile += fmt.Sprint("\nSecond hand score:", p.SplitHand.Cards.BlackjackScore())
	}

	return profile
}

func (p *Player) Hit(shoe *deck.Deck, hand *PlayerHand) deck.Card {
	hitCard := (*shoe).PopCard()
	hand.Cards = append(hand.Cards, hitCard)
	return hitCard
}
