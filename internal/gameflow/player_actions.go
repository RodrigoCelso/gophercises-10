package gameflow

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/deck"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

func playerPlay(match *game.Game, p *bjackplayer.Player) {
	currentHand := &p.Hand
	splitOptionTxt := "4. Split\n"
	doubleDownOptionTxt := "3. Double down\n"
	if (*currentHand)[0].Value != (*currentHand)[1].Value {
		splitOptionTxt = ""
	}
	if p.Bet*2 > p.Chips {
		doubleDownOptionTxt = ""
	}
	fmt.Println("=====================\n", p, "\n=====================")
	for {
		fmt.Printf("Would you like to Hit or Stand? (Current hand score: %d)\n", bjackplayer.Score(*currentHand))
		fmt.Println("1. Hit\n2. Stand\n" + doubleDownOptionTxt + splitOptionTxt + "5. Check table")

		var playerChoice int
		fmt.Scanf("%d\n", &playerChoice)

		if playerChoice == 1 {
			cardHitted := p.Hit(&match.Shoe)
			fmt.Print("\nYou hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n\n")
			if bjackplayer.Score(p.Hand) > 21 {
				fmt.Print("BUSTED! ", bjackplayer.Score(p.Hand), "\n")
				p.Hand = []deck.Card{}
				break
			}
			pScore := bjackplayer.Score(p.Hand)
			if pScore > match.PlayerMaxScore {
				match.PlayerMaxScore = pScore
			}
		} else if playerChoice == 2 {
			if p.SplitHand != nil && currentHand != &p.SplitHand {
				currentHand = &p.SplitHand
				continue
			}
			break
		} else if playerChoice == 3 {
			// Double down
			if doubleDownOptionTxt != "" {
				fmt.Println("DOUBLE DOWN!")
				p.Bet *= 2
				cardHitted := p.Hit(&match.Shoe)
				fmt.Print("\nYou hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n\n")
				if bjackplayer.Score(p.Hand) > 21 {
					fmt.Print("BUSTED! ", bjackplayer.Score(p.Hand), "\n")
					p.Hand = []deck.Card{}
				}
				if p.SplitHand != nil && currentHand != &p.SplitHand {
					currentHand = &p.SplitHand
					continue
				}
				break
			}
		} else if playerChoice == 4 {
			// Split
			if splitOptionTxt != "" {
				handSize := len(p.Hand)
				p.SplitHand = p.Hand[handSize>>1:]
				p.Hand = p.Hand[:handSize>>1]
			}
		} else if playerChoice == 5 {
			checkTable(match, p)
		} else {
			// invalid command
			fmt.Println("Invalid command, please try again.")
			continue
		}
	}
	fmt.Print("Final Score: ", bjackplayer.Score(p.Hand), "\n=====================\n")
}

func checkTable(match *game.Game, pCall *bjackplayer.Player) {
	fmt.Println("=====================")
	fmt.Println(match.Dealer)
	fmt.Println("=====================")
	fmt.Println("\n=====================")
	for _, p := range match.Players {
		fmt.Println(p)
		if p == pCall {
			fmt.Println("======= You ^ =======")
			continue
		}
		fmt.Println("=====================")
	}
}
