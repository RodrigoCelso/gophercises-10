package player

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/deck"
)

type PlayerState uint8

const (
	Default PlayerState = iota
	Blackjack
)

type Player struct {
	Name  string
	Hand  deck.Deck
	Chips int
	Bet   int
	State PlayerState
}

func New(name string) *Player {
	return &Player{Name: name}
}

func (p *Player) Score() int {
	var aces uint8
	var scoreValue int
	for _, card := range p.Hand {
		cardValue, isAce := card.BlackjackScoreWithAce()
		scoreValue += cardValue
		if isAce {
			aces++
		}
	}
	for range aces {
		if scoreValue > 21 {
			scoreValue -= 10
			aces--
		} else {
			aces--
		}
	}
	return scoreValue
}

func (p *Player) String() string {
	profile := fmt.Sprint(p.Name, ": Cards\n")
	for _, card := range p.Hand {
		profile += fmt.Sprintln("-", card)
	}
	profile += fmt.Sprintln("Score:", p.Score())
	return profile
}

func (p *Player) Hit(shoe *deck.Deck) deck.Card {
	shoeLast := len(*shoe) - 1
	hitCard := (*shoe)[shoeLast]
	*shoe = (*shoe)[:shoeLast]
	p.Hand = append(p.Hand, hitCard)
	return hitCard
}
