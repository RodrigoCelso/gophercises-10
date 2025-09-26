package gameflow

import (
	"fmt"
	"strconv"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/controller"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

func saveScore(playerName string, playerChips int, playerBet int) error {
	bucketKey := playerName + " - " + time.Now().Format("02/01/2006 15:04:05")
	var bucketValue string
	if playerBet == 0 {
		bucketValue = "Tie - Total: " + strconv.Itoa(playerChips) + " chips"
	} else {
		err := controller.InsertRemoveChips(playerName, playerBet)
		if err != nil {
			return fmt.Errorf("couldn't update player chips: %w", err)
		}
		bucketValue = "Lose - paid " + strconv.Itoa(playerBet) + " Chips - Total: " + strconv.Itoa(playerChips+playerBet) + " chips"
	}

	err := controller.NewScoreboardEntry(bucketKey, bucketValue)
	if err != nil {
		return fmt.Errorf("couldn't write into database: %w", err)
	}
	return nil
}

func settleUsers(dealer *game.Player, users []*game.Player) error {
	dealerScore := dealer.MainHand.Cards.BlackjackScore()
	for _, u := range users {
		currentHand := &u.MainHand
		for {
			uScore := u.MainHand.Cards.BlackjackScore()

			if uScore == dealerScore && u.MainHand.NaturalBlackjack == dealer.MainHand.NaturalBlackjack {
				// tied
				fmt.Printf("%s and %s tied (user %dx%d dealer)\n", u.Name, dealer.Name, uScore, dealerScore)
				err := saveScore(u.Name, u.Chips, 0)
				if err != nil {
					return fmt.Errorf("couldn't save the score: %w", err)
				}
			} else if dealerScore > uScore || dealer.MainHand.NaturalBlackjack && !currentHand.NaturalBlackjack {
				// player lost
				fmt.Printf("%s won against %s (dealer %dx%d user)\n", dealer.Name, u.Name, dealerScore, uScore)
				err := saveScore(u.Name, u.Chips, -currentHand.Bet)
				if err != nil {
					return fmt.Errorf("couldn't save the score: %w", err)
				}
			} else {
				// player won
				fmt.Printf("%s won against %s (user %dx%d dealer)\n", u.Name, dealer.Name, uScore, dealerScore)
				err := saveScore(u.Name, u.Chips, currentHand.Bet)
				if err != nil {
					return fmt.Errorf("couldn't save the score: %w", err)
				}
			}

			if u.Splitted {
				if currentHand == &u.SplitHand {
					break
				}
				currentHand = &u.SplitHand
			} else {
				break
			}
		}
	}
	return nil
}

func settleNPCs(dealer *game.Player, npcs []*game.NPCPlayer) {
	dealerScore := dealer.MainHand.Cards.BlackjackScore()
	for _, n := range npcs {
		nScore := n.MainHand.Cards.BlackjackScore()
		if nScore == dealerScore && n.MainHand.NaturalBlackjack == dealer.MainHand.NaturalBlackjack {
			// tie
			fmt.Printf("%s and %s tied (player %dx%d dealer)\n", n.Name, dealer.Name, nScore, dealerScore)
			continue
		}

		if dealerScore > nScore || dealer.MainHand.NaturalBlackjack && !n.MainHand.NaturalBlackjack {
			// npc lost
			fmt.Printf("%s won against %s (dealer %dx%d player)\n", dealer.Name, n.Name, dealerScore, nScore)
			continue
		}

		// npc won
		fmt.Printf("%s won against %s (player %dx%d dealer)\n", n.Name, dealer.Name, nScore, dealerScore)
	}
}

func settleWinner(users []*game.Player, npcs []*game.NPCPlayer, dealer *game.Player) {
	settleUsers(dealer, users)
	settleNPCs(dealer, npcs)
}
