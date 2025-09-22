package gameflow

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/deck"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

// A Inteligência para este jogo está em como reagir as cartas do dealer
// Primeira carta do dealer é um número entre 2 a 6 = bom
// primeira carta do dealer é um número entre 7 a 9 = médio
// primeira carta do dealer é um ás, J, Q, K ou 10 = ruim
const DEALER_SCORE_INFLUENCE float32 = 0.1

func NPCPlay(match *game.Game, n *bjackplayer.Player) {
	dealerScore := match.Dealer.MainHand.Cards[0].BlackjackScore()
	fmt.Printf("=====================\n")
	for {
		handScore := n.MainHand.Cards.BlackjackScore()
		if handScore > 21 {
			fmt.Println("Busted!", handScore)
			fmt.Printf("=====================\n")
			break
		}

		// Calculates risk (starts at 11)
		handRisk := float64(handScore-11) / 10.0

		// Calculates dealer first card threat
		dealerThreat := 1.0
		if dealerScore < 5 {
			dealerThreat = 0.5
		}
		if dealerScore >= 10 {
			dealerThreat = 1.5
			if match.Dealer.MainHand.Cards[0].Value != 1<<9 {
				dealerThreat *= 2
			}
		}

		// Calculates fear based on intelligence level and the cards in the table
		fear := handRisk * dealerThreat * match.NPCIntelligence

		// If there is no fear, then hit
		if fear <= 0 {
			n.Hit(&match.Shoe, &n.MainHand)
			handScore = n.MainHand.Cards.BlackjackScore()
			if handScore > 21 {
				n.MainHand.Cards = deck.Deck{}
				fmt.Println(n.Name, "- I'm BUSTED :(")
				fmt.Printf("=====================\n")
				break
			}
			fmt.Printf("%s - Hit: %d\n", n.Name, handScore)
			time.Sleep(FOCUS_TIME)
		} else {
			// If the score is already 21, stand
			if handScore == 21 {
				fmt.Printf("%s - Stand: %d\n\n", n.Name, handScore)
				fmt.Printf("=====================\n")
				break
			}

			// Generates random behaviour to simulate courage
			courage := rand.Float64()
			if courage > fear {
				n.Hit(&match.Shoe, &n.MainHand)
				handScore = n.MainHand.Cards.BlackjackScore()
				if handScore > 21 {
					n.MainHand.Cards = deck.Deck{}
					fmt.Println(n.Name, "- I'm BUSTED :(")
					fmt.Printf("=====================\n")
					break
				}
				fmt.Printf("%s - Hit: %d\n", n.Name, handScore)
				time.Sleep(FOCUS_TIME)
			}
			fmt.Printf("%s - Stand: %d\n", n.Name, handScore)
			fmt.Printf("=====================\n")
			break
		}
	}
}
