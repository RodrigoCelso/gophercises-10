package gameflow

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/deck"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

func NPCPlay(match *game.Game, n *game.NPCPlayer) {
	dealerScore := match.Dealer.MainHand.Cards[0].BlackjackScore()
	fmt.Printf("=====================\n")
	fmt.Println(n.Name, "- First Score:", n.MainHand.Cards.BlackjackScore())
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
		fear := handRisk * dealerThreat * float64(match.NPCIntelligence)

		// If there is no fear, then hit
		if fear <= 0 {
			n.MainHand.Hit(match)
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
				fmt.Printf("%s - Stand: %d\n", n.Name, handScore)
				fmt.Printf("=====================\n")
				break
			}

			// Generates random behaviour to simulate courage
			courage := rand.Float64()
			if courage > fear {
				n.MainHand.Hit(match)
				handScore = n.MainHand.Cards.BlackjackScore()
				if handScore > 21 {
					n.MainHand.Cards = deck.Deck{}
					fmt.Println(n.Name, "- I'm BUSTED :(")
					fmt.Printf("=====================\n")
					break
				}
				fmt.Printf("%s - Hit: %d\n", n.Name, handScore)
			}
			fmt.Printf("%s - Stand: %d\n", n.Name, handScore)
			fmt.Printf("=====================\n")
			break
		}
	}
}

func NPCTricksterPlay(match *game.Game, npc *game.NPCPlayer) {
	handScore := npc.MainHand.Cards.BlackjackScore()
	fmt.Printf("=====================\n")
	fmt.Println(npc.Name, "- First score:", handScore)
	if handScore == 21 {
		npc.MainHand.NaturalBlackjack = true
		fmt.Println(npc.Name, "- Natural Blackjack:", handScore)
		return
	}
	for {
		counterProbability := float64(match.CardCounter) / float64(match.ShoeDecks)

		riskSigmoid := 1.0 / (1.0 + math.Exp(counterProbability))
		handRisk := (float64(handScore-11) / 10.0) * riskSigmoid

		if counterProbability == 0 && handScore < 17 {
			// hit
			npc.MainHand.Hit(match)
			handScore = npc.MainHand.Cards.BlackjackScore()

			if handScore > 21 {
				npc.MainHand.Cards = deck.Deck{}
				fmt.Println(npc.Name, "- I'm BUSTED :(")
				fmt.Printf("=====================\n")
				break
			}

			fmt.Println(npc.Name, "- Hit:", handScore)
			continue
		}

		if handRisk < 0.1 {
			// hit
			npc.MainHand.Hit(match)
			handScore = npc.MainHand.Cards.BlackjackScore()

			if handScore > 21 {
				npc.MainHand.Cards = deck.Deck{}
				fmt.Println(npc.Name, "- I'm BUSTED :(")
				fmt.Printf("=====================\n")
				break
			}

			fmt.Println(npc.Name, "- Hit:", handScore)
			continue
		}
		fmt.Println(npc.Name, "- Stand:", handScore)
		break
	}
}
