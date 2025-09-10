package gameflow

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/deck"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

func playerPlay(round *game.Game, p *bjackplayer.Player) {
	fmt.Println("=====================\n", p, "\n=====================")
	for {
		fmt.Printf("Would you like to Hit or Stand? (Current Score: %d)\n", p.Score())
		fmt.Println("1. Hit\n2. Stand\n3. Check table")

		var playerChoice int
		fmt.Scanf("%d\n", &playerChoice)

		if playerChoice == 1 {
			cardHitted := p.Hit(&round.Shoe)
			fmt.Print("\nYou hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n\n")
			if p.Score() > 21 {
				fmt.Print("BUSTED! ", p.Score(), "\n")
				p.Hand = []deck.Card{}
				break
			}
			pScore := p.Score()
			if pScore > round.PlayerMaxScore {
				round.PlayerMaxScore = pScore
			}
		} else if playerChoice == 2 {
			break
		} else if playerChoice == 3 {
			checkTable(round, p)
		} else {
			// invalid command
			fmt.Println("Invalid command, please try again.")
			continue
		}
	}
	fmt.Print("Final Score: ", p.Score(), "\n=====================\n")
}

func checkTable(round *game.Game, pCall *bjackplayer.Player) {
	fmt.Println("=====================")
	fmt.Println(round.Dealer)
	fmt.Println("=====================")
	fmt.Println("\n=====================")
	for _, p := range round.Players {
		fmt.Println(p)
		if p == pCall {
			fmt.Println("======= You ^ =======")
			continue
		}
		fmt.Println("=====================")
	}
}
