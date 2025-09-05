package main

import (
	"fmt"
	"time"

	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/game"
	"github.com/RodrigoCelso/gophercises-10/player"
)

const FOCUS_TIME = 2 * time.Second

var playerMaxScore int

func main() {
	// New round created
	newRound := game.New(1)
	playerMaxScore = 0

	// Deal cards
	fmt.Print("Welcome to the blackjack game, i'm your dealer\ndealing cards...\n\n")
	time.Sleep(FOCUS_TIME)
	dealCards(newRound.Players, newRound.Dealer, &newRound.Shoe)

	// Dealer reveals first card
	fmt.Println("this is my first card:")
	dealerFlipCard(newRound.Dealer, 1)
	time.Sleep(FOCUS_TIME)

	// Players turn
	fmt.Println("this is your hand:")
	for _, p := range newRound.Players {
		fmt.Println(p)
		time.Sleep(FOCUS_TIME)
	}
	playersPlay(newRound.Players, &newRound.Shoe)
	time.Sleep(FOCUS_TIME)

	// Dealer reveals second card
	fmt.Println("this is my second card:")
	dealerFlipCard(newRound.Dealer, 2)
	fmt.Println("my current hand is:\n", newRound.Dealer)
	time.Sleep(FOCUS_TIME)

	// Dealer turn
	dealerPlay(newRound.Dealer, &newRound.Shoe)
	time.Sleep(FOCUS_TIME)

	// Decide winners
	compareCards(newRound.Players, newRound.Dealer)
}

func dealCards(players []*player.Player, dealer *player.Player, shoe *deck.Deck) {
	for _, p := range players {
		p.Hit(shoe)
		p.Hit(shoe)
		pScore := p.Score()
		if pScore == 21 {
			// Natural Blackjack
			p.State = player.Blackjack
			playerMaxScore = pScore
		}
	}
	dealer.Hit(shoe)
	dealer.Hit(shoe)
	if dealer.Score() == 21 {
		// Natural Blackjack
		dealer.State = player.Blackjack
	}
}

func dealerFlipCard(dealer *player.Player, cardNumber int) {
	c := dealer.Hand[cardNumber-1]
	fmt.Print(c, " - Value: ", c.BlackjackScore(), "\n\n")
}

func playersPlay(players []*player.Player, shoe *deck.Deck) {
	for _, p := range players {
		for {
			fmt.Println(p.Name, "your current score is:", p.Score())
			fmt.Println("Would you like to Hit or Stand?")
			fmt.Println("0. Your cards\n1. Hit\n2. Stand")

			playerChoice := -1
			fmt.Scanf("%d\n", &playerChoice)

			if playerChoice == 0 {
				fmt.Println(p.Hand)
				continue
			} else if playerChoice == 1 {
				cardHitted := p.Hit(shoe)
				fmt.Print("\nYou hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n\n")
				if p.Score() > 21 {
					fmt.Print("BUSTED! ", p.Score(), "\n")
					p.Hand = []deck.Card{}
					break
				}
				pScore := p.Score()
				if pScore > playerMaxScore {
					playerMaxScore = pScore
				}
			} else if playerChoice == 2 {
				break
			} else {
				// invalid command
				fmt.Println("Invalid command, please try again.")
				continue
			}
		}
		fmt.Print("Final Score: ", p.Score(), "\n\n")
	}
}

func dealerPlay(dealer *player.Player, shoe *deck.Deck) {
	score := dealer.Score()
	for score < 17 || score < playerMaxScore {
		cardHitted := dealer.Hit(shoe)
		score = dealer.Score()
		fmt.Print("\nI hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n")
		fmt.Print("My current score: ", score, "\n\n")
		time.Sleep(FOCUS_TIME)
		if score > playerMaxScore {
			break
		}
	}
	if score > 21 {
		fmt.Print("Dealer BUSTED ", dealer.Score(), "\n")
		dealer.Hand = []deck.Card{}
	}
}

func compareCards(players []*player.Player, dealer *player.Player) {
	dealerScore := dealer.Score()
	for _, p := range players {
		pScore := p.Score()
		if (pScore == dealerScore) || (p.State == dealer.State && !(p.State == player.Default && dealer.State == player.Default)) {
			// tie
			fmt.Println(dealer.Name, "and", p.Name, "tied")
			return
		}

		if dealer.State == player.Blackjack || dealerScore > pScore {
			// dealer won
			fmt.Println(dealer.Name, "won against", p.Name)
			return
		}

		// dealer lost
		fmt.Println(p.Name, "won against", dealer.Name)
	}
}
