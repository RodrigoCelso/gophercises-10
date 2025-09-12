package gameflow

import (
	"fmt"
	"strconv"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/controller"
)

func settleWinner(players []*bjackplayer.Player, dealer *bjackplayer.Player) {
	dealerScore := bjackplayer.Score(dealer.Hand)
	for _, p := range players {
		pScore := bjackplayer.Score(p.Hand)
		if (pScore == dealerScore) || (p.State == dealer.State && !(p.State == bjackplayer.Default && dealer.State == bjackplayer.Default)) {
			// tie
			fmt.Println(dealer.Name, "and", p.Name, "tied")
			if p.Playable {
				bucketKey := p.Name + " - " + time.Now().Format("02/01/2006 15:04:05")
				bucketValue := "Tie - Total: " + strconv.Itoa(p.Chips) + " chips"

				err := controller.NewScoreboardEntry(bucketKey, bucketValue)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}
			continue
		}

		if dealer.State == bjackplayer.Blackjack || dealerScore > pScore {
			// dealer won
			fmt.Println(dealer.Name, "won against", p.Name)
			if p.Playable {
				err := controller.InsertChips(p.Name, -p.Bet)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				bucketKey := p.Name + " - " + time.Now().Format("02/01/2006 15:04:05")
				bucketValue := "Lose - paid " + strconv.Itoa(p.Bet) + " Chips - Total: " + strconv.Itoa(p.Chips-p.Bet) + " chips"

				err = controller.NewScoreboardEntry(bucketKey, bucketValue)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
			}
			continue
		}

		// dealer lost
		fmt.Println(p.Name, "won against", dealer.Name)
		if p.Playable {
			err := controller.InsertChips(p.Name, p.Bet)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			bucketKey := p.Name + " - " + time.Now().Format("02/01/2006 15:04:05")
			bucketValue := "Win - Gained " + strconv.Itoa(p.Bet) + " Chips - Total: " + strconv.Itoa(p.Chips+p.Bet) + " chips"

			err = controller.NewScoreboardEntry(bucketKey, bucketValue)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

	}
}
