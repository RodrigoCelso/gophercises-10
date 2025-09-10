package gameflow

import (
	"fmt"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

const FOCUS_TIME = 2 * time.Second

var round *game.Game

func StartGame() {
	MainMenu()
}

func MainMenu() {
	for {
		fmt.Println("Blackjack\n0. Exit\n1. Start new Game?\n2. Scoreboard")
		var menuChoice uint8
		fmt.Scanf("%d\n", &menuChoice)
		switch menuChoice {
		case 0:
			return
		case 1:
			fmt.Println("How many players?")
			var playersQuantityChoice uint8
			fmt.Scanf("%d\n", &playersQuantityChoice)
			round = game.New(int(playersQuantityChoice))
			NewGame()
		case 2:
			Scoreboard()
		default:
			continue
		}
	}
}

func NewGame() {
	fmt.Println("Make your bets")
	for _, p := range round.Players {
		var pBet int
		fmt.Print(p.Name, " - ")
		fmt.Scanf("%d\n", &pBet)
		p.Bet = pBet
	}

	// Deal cards
	fmt.Print("Welcome to the blackjack game, i'm your dealer\ndealing cards...\n\n")
	time.Sleep(FOCUS_TIME)
	DealCards(round)

	// Dealer reveals first card
	fmt.Println("this is my first card:")
	fmt.Print(FlipCard(round.Dealer, 1), "\n\n")
	time.Sleep(FOCUS_TIME)

	// Players turn
	for _, p := range round.Players {
		playerPlay(round, p)
		time.Sleep(FOCUS_TIME)
	}

	// Dealer reveals second card
	fmt.Println("this is my second card:")
	fmt.Print(FlipCard(round.Dealer, 2), "\n\n")
	fmt.Println("my current hand is:\n", round.Dealer)
	time.Sleep(FOCUS_TIME)

	// Dealer turn
	DealerPlay(round)
	time.Sleep(FOCUS_TIME)

	// Decide winners
	CompareCards(round.Players, round.Dealer)
}
