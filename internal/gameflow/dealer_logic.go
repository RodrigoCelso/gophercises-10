package gameflow

import (
	"fmt"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/deck"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

func DealCards(round *game.Game) {
	for _, p := range round.Players {
		p.Hit(&round.Shoe)
		p.Hit(&round.Shoe)
		pScore := p.Score()
		if pScore == 21 {
			// Natural Blackjack
			p.State = bjackplayer.Blackjack
		}
		if pScore > round.PlayerMaxScore {
			round.PlayerMaxScore = pScore
		}
	}
	round.Dealer.Hit(&round.Shoe)
	round.Dealer.Hit(&round.Shoe)
	if round.Dealer.Score() == 21 {
		// Natural Blackjack
		round.Dealer.State = bjackplayer.Blackjack
	}
}

func FlipCard(dealer *bjackplayer.Player, cardNumber int) *deck.Card {
	return &dealer.Hand[cardNumber-1]
}

func DealerPlay(round *game.Game) {
	score := round.Dealer.Score()
	for score < 17 || (score < round.PlayerMaxScore && len(round.Players) == 1) {
		cardHitted := round.Dealer.Hit(&round.Shoe)
		score = round.Dealer.Score()
		fmt.Print("\nI hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n")
		fmt.Print("My current score: ", score, "\n\n")
		time.Sleep(time.Second)
		if score > round.PlayerMaxScore {
			break
		}
	}
	if score > 21 {
		fmt.Print("Dealer BUSTED ", round.Dealer.Score(), "\n")
		round.Dealer.Hand = []deck.Card{}
	}
}
