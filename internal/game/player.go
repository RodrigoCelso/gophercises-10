package game

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

func NewPlayer(name string) *Player {
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

func (h *PlayerHand) Hit(match *Game) deck.Card {
	hitCard := match.Shoe.PopCard()
	match.CardCounter += cardCount(hitCard)
	h.Cards = append(h.Cards, hitCard)
	return hitCard
}

func cardCount(card deck.Card) int {
	switch card.BlackjackScore() {
	case 2, 3, 4, 5, 6:
		return 1
	case 7, 8, 9:
		return 0
	case 10, 11:
		return -1
	}
	return 0
}
