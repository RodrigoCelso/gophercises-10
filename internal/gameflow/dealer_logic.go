package gameflow

import (
	"fmt"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/deck"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

func dealCards(match *game.Game) {
	// Deal users cards
	for _, u := range match.Users {
		u.MainHand.Hit(match)
		u.MainHand.Hit(match)
		uScore := u.MainHand.Cards.BlackjackScore()
		if uScore == 21 {
			// Natural Blackjack
			u.MainHand.NaturalBlackjack = true
		}
		// Update Player Max Score except for natural blackjacks
		if uScore > match.PlayerMaxScore && !u.MainHand.NaturalBlackjack {
			match.PlayerMaxScore = uScore
		}
	}

	// Deal NPCs cards
	for _, n := range match.NPCs {
		n.MainHand.Hit(match)
		n.MainHand.Hit(match)
		nScore := n.MainHand.Cards.BlackjackScore()
		if nScore == 21 {
			n.MainHand.NaturalBlackjack = true
		}
		// Update Player Max Score except for natural blackjacks
		if nScore > match.PlayerMaxScore && !n.MainHand.NaturalBlackjack {
			match.PlayerMaxScore = nScore
		}
	}

	dealer := match.Dealer
	// Deal dealer cards
	dealer.MainHand.Hit(match)
	dealer.MainHand.Hit(match)
	if dealer.MainHand.Cards.BlackjackScore() == 21 {
		// Natural Blackjack
		dealer.MainHand.NaturalBlackjack = true
	}
}

func flipCard(dealer *game.Player, cardNumber int) *deck.Card {
	return &dealer.MainHand.Cards[cardNumber-1]
}

func dealerPlay(match *game.Game) {
	dealer := match.Dealer
	score := dealer.MainHand.Cards.BlackjackScore()
	for score < 17 || (score < match.PlayerMaxScore && len(match.Users)+len(match.NPCs) == 1) {
		cardHitted := dealer.MainHand.Hit(match)
		score = dealer.MainHand.Cards.BlackjackScore()
		fmt.Print("\nI hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n")
		fmt.Print("My current score: ", score, "\n\n")
		time.Sleep(time.Second)
		if score > match.PlayerMaxScore {
			break
		}
	}
	if score > 21 {
		fmt.Print("Dealer BUSTED ", dealer.MainHand.Cards.BlackjackScore(), "\n")
		dealer.MainHand.Cards = []deck.Card{}
	}
}
