package gameflow

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/internal/deck"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

func playerPlay(match *game.Game, p *game.Player) {
	currentHand := &p.MainHand
	splitOptionTxt := "4. Split\n"
	doubleDownOptionTxt := "3. Double down\n"
	if currentHand.Cards[0].Value != currentHand.Cards[1].Value {
		splitOptionTxt = ""
	}
	if currentHand.Bet*2 > p.Chips || currentHand.Bet == 0 {
		doubleDownOptionTxt = ""
	}
	fmt.Println("=====================\n", p, "\n=====================")
	for {
		fmt.Printf("Would you like to Hit or Stand? (Current hand score: %d)\n", currentHand.Cards.BlackjackScore())
		fmt.Println("1. Hit\n2. Stand\n" + doubleDownOptionTxt + splitOptionTxt + "5. Check table")

		var playerChoice int
		fmt.Scanf("%d\n", &playerChoice)

		if playerChoice == 1 {
			cardHitted := currentHand.Hit(match)
			fmt.Print("\nYou hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n\n")
			if currentHand.Cards.BlackjackScore() > 21 {
				fmt.Print("BUSTED! ", currentHand.Cards.BlackjackScore(), "\n")
				p.MainHand.Cards = []deck.Card{}
				if p.SplitHand.Cards != nil && currentHand != &p.SplitHand {
					currentHand = &p.SplitHand
					continue
				}
				break
			}
			pScore := currentHand.Cards.BlackjackScore()
			if pScore > match.PlayerMaxScore {
				match.PlayerMaxScore = pScore
			}
		} else if playerChoice == 2 {
			if p.SplitHand.Cards != nil && currentHand != &p.SplitHand {
				currentHand = &p.SplitHand
				continue
			}
			break
		} else if playerChoice == 3 {
			// Double down
			if doubleDownOptionTxt != "" {
				fmt.Println("DOUBLE DOWN!")
				currentHand.Bet *= 2
				cardHitted := currentHand.Hit(match)
				fmt.Print("\nYou hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n\n")
				if currentHand.Cards.BlackjackScore() > 21 {
					fmt.Print("BUSTED! ", currentHand.Cards.BlackjackScore(), "\n")
					p.MainHand.Cards = []deck.Card{}
				}
				if p.SplitHand.Cards != nil && currentHand != &p.SplitHand {
					currentHand = &p.SplitHand
					continue
				}
				break
			}
		} else if playerChoice == 4 {
			// Split
			if splitOptionTxt != "" {
				splitOptionTxt = ""
				handSize := len(p.MainHand.Cards)
				p.SplitHand.Bet = p.MainHand.Bet >> 1
				p.MainHand.Bet -= p.SplitHand.Bet
				p.SplitHand.Cards = p.MainHand.Cards[handSize>>1:]
				p.MainHand.Cards = p.MainHand.Cards[:handSize>>1]
			}
		} else if playerChoice == 5 {
			checkTable(match, p)
		} else {
			// invalid command
			fmt.Println("Invalid command, please try again.")
			continue
		}
	}
	fmt.Print("Final Score: ", currentHand.Cards.BlackjackScore(), "\n=====================\n")
}

func checkTable(match *game.Game, pCall *game.Player) {
	fmt.Println("=====================")
	fmt.Println(match.Dealer)
	fmt.Println("=====================")
	fmt.Println("\n=====================")
	for _, p := range match.Users {
		fmt.Println(p)
		if p == pCall {
			fmt.Println("======= You ^ =======")
			continue
		}
		fmt.Println("=====================")
	}
}
