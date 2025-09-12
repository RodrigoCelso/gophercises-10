package bjackplayer

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/internal/deck"
)

type PlayerState uint8

const (
	Default PlayerState = iota
	Blackjack
)

type Player struct {
	Name      string
	Hand      deck.Deck
	SplitHand deck.Deck
	Chips     int
	Bet       int
	State     PlayerState
	Playable  bool
}

func Score(hand deck.Deck) int {
	var aces uint8
	var scoreValue int
	for _, card := range hand {
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
	profile := fmt.Sprint(p.Name, ":\nHand cards\n")
	for _, card := range p.Hand {
		profile += fmt.Sprintln("-", card)
	}
	if p.SplitHand != nil {
		profile += fmt.Sprintln("\nSecond hand cards")
		for _, card := range p.SplitHand {
			profile += fmt.Sprintln("-", card)
		}
	}
	profile += fmt.Sprint("Hand score:", Score(p.Hand))

	if p.SplitHand != nil {
		profile += fmt.Sprint("\nSecond hand score:", Score(p.SplitHand))
	}

	return profile
}

func (p *Player) Hit(shoe *deck.Deck) deck.Card {
	shoeLast := len(*shoe) - 1
	hitCard := (*shoe)[shoeLast]
	*shoe = (*shoe)[:shoeLast]
	p.Hand = append(p.Hand, hitCard)
	return hitCard
}
