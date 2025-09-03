package player

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/deck"
)

type IPlayer interface {
	Hit()
	Stand()
	DoubleDown()
	Split()
}

type Player struct {
	Name  string
	Hand  deck.Deck
	Chips int
	Bet   int
}

func New(name string) *Player {
	return &Player{Name: name}
}

func NewDealer() *Player {
	return &Player{Name: "Dealer"}
}

func (p *Player) Score() int {
	var aces uint8
	var scoreValue int
	for _, card := range p.Hand {
		cardValue, isAce := card.BJScore()
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
	situation := fmt.Sprintln(p.Name, "- Score: ", p.Score(), "\nCards:")
	for _, card := range p.Hand {
		situation += fmt.Sprintln(card)
	}
	return situation
}
