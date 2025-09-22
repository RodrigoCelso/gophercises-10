package gameflow

import (
	"fmt"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/controller"
	"github.com/RodrigoCelso/gophercises-10/internal/game"
)

const FOCUS_TIME = 2 * time.Second

var match *game.Game

func StartGame() {
	MainMenu()
}

func MainMenu() {
	for {
		fmt.Println("Blackjack\n0. Exit\n1. Start new Game?\n2. Insert user chips\n3. Scoreboard\n4. Check player chips")
		var menuChoice uint8
		fmt.Scanf("%d\n", &menuChoice)
		switch menuChoice {
		case 0:
			return
		case 1:
			// NPC quantity
			fmt.Println("How many players (non playable)?")
			var npcQuantityChoice uint8
			fmt.Scanln(&npcQuantityChoice)

			// User quantity
			fmt.Println("How many users (playable)?")
			var playersQuantityChoice uint8
			fmt.Scanln(&playersQuantityChoice)
			var users []*bjackplayer.Player

			// Users Login
			for range playersQuantityChoice {
				var playerName string
				var playerChips int

				fmt.Print("Name: ")
				fmt.Scanln(&playerName)

				playerChips, err := controller.GetChips(playerName)
				if err != nil {
					fmt.Println("Error:", err)
				}
				users = append(users, &bjackplayer.Player{Name: playerName, Chips: playerChips})
			}

			// Create match
			match = game.New(int(npcQuantityChoice), users)
			NewGame()
		case 2:
			fmt.Print("Player name: ")
			var userName string
			fmt.Scanln(&userName)

			fmt.Print("How many chips? ")
			var playerChips int
			fmt.Scanln(&playerChips)

			err := controller.InsertRemoveChips(userName, playerChips)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case 3:
			scoreboardMap, err := controller.Scoreboard()
			if err != nil {
				fmt.Println("Error:", err)
			}
			for k, v := range scoreboardMap {
				fmt.Println(k, ":", v)
			}
		case 4:
			fmt.Print("User name: ")
			var userName string
			fmt.Scanln(&userName)
			chips, err := controller.GetChips(userName)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Print(chips, " chips\n\n")
		default:
			continue
		}
	}
}

func NewGame() {
	// Users bet
	for _, p := range match.Users {
		if p.Chips > 0 {
			fmt.Printf("Make your bet %s: (chips available: %d)\n", p.Name, p.Chips)
			var pBet int

			for {
				fmt.Print(p.Name, " - ")
				fmt.Scanf("%d\n", &pBet)

				if pBet < 0 {
					fmt.Println("You can't bet negative chips")
					continue
				}
				if pBet > p.Chips {
					fmt.Println("You don't have enough chips")
					continue
				}
				break
			}
			p.MainHand.Bet = pBet
		}
	}

	// Deal cards
	fmt.Print("Welcome to the blackjack game, i'm your dealer\ndealing cards...\n\n")
	time.Sleep(FOCUS_TIME)
	dealCards(match)

	// Dealer reveals first card
	fmt.Println("this is my first card:")
	dealerFirst := flipCard(match.Dealer, 1)
	fmt.Print(dealerFirst, "\n\n")
	time.Sleep(FOCUS_TIME)

	// Players turn
	for _, u := range match.Users {
		playerPlay(match, u)
		time.Sleep(FOCUS_TIME)
	}

	// NPCs turn
	for _, n := range match.NPCs {
		NPCPlay(match, n)
		time.Sleep(FOCUS_TIME)
	}

	// Dealer reveals second card
	fmt.Println("this is all my current cards:")
	fmt.Print(match.Dealer, "\n\n")
	time.Sleep(FOCUS_TIME)

	// Dealer turn
	dealerPlay(match)
	time.Sleep(FOCUS_TIME)

	// Decide winners
	settleWinner(match.Users, match.NPCs, match.Dealer)
}
