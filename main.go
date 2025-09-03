package main

import (
	"fmt"
	"time"

	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/game"
	"github.com/RodrigoCelso/gophercises-10/player"
)

const FOCUS_TIME = 2 * time.Second

func main() {
	game := initGame()
	round(&game)
}

func initGame() game.Game {
	game := game.Game{Shoe: *deck.New(deck.WithMultipleDeckSize(4), deck.WithShuffle()), Dealer: player.NewDealer(), Players: []*player.Player{player.New("Player")}}
	game.StartRound()

	dealerIntro(&game)
	return game
}

func dealerIntro(game *game.Game) {
	fmt.Println(game.Dealer.Name, "- My first card:", game.Dealer.Hand[0])
	pScore, _ := game.Dealer.Hand[0].BJScore()
	fmt.Print("Partial score: ", pScore, "\n\n")
}

func userPlay(game *game.Game, user *player.Player) int {
	var userChoice int
	score := user.Score()
	fmt.Println(user)

	fmt.Println("What are you going to do?")
	for {
		fmt.Println("1. Hit")
		fmt.Println("2. Stand")
		fmt.Scan(&userChoice)

		if userChoice == 1 {
			card := game.Hit(user)
			score = user.Score()
			if score > 21 {
				fmt.Println("Card hitted:", card)
				break
			}
			fmt.Println("Card hitted:", card)
			fmt.Println("Current Score:", score)
			time.Sleep(FOCUS_TIME)
			continue
		}

		if userChoice == 2 {
			fmt.Print("Stand: ", score, "\n\n")
			time.Sleep(FOCUS_TIME)
			return score
		}

		fmt.Println("Invalid action")
		return score
	}

	fmt.Print(user.Name, " - Score: ", score, "\n\n")
	time.Sleep(FOCUS_TIME)
	return score
}

func dealerPlay(game *game.Game, dealer *player.Player) int {
	score := dealer.Score()
	fmt.Println(dealer)
	fmt.Println("Score: ", score)
	time.Sleep(2 * time.Second)
	for score <= 16 {
		newCard := game.Hit(dealer)
		score = dealer.Score()
		fmt.Println("Hit:", newCard, "\nScore:", score)
		time.Sleep(FOCUS_TIME)
	}
	fmt.Println("Stand:", score)
	return score
}

func round(game *game.Game) {
	scorePlayers := make(map[string]int)
	for playersLen := range game.Players {
		scorePlayers[game.Players[playersLen].Name] = userPlay(game, game.Players[playersLen])
	}
	scoreDealer := dealerPlay(game, game.Dealer)
	for name, score := range scorePlayers {
		if scoreDealer > score {
			fmt.Println(game.Dealer.Name, "won against", name)
			return
		}
		if scoreDealer == score {
			fmt.Println(game.Dealer.Name, "and", name, "tied")
			return
		}
		if scoreDealer < score {
			fmt.Println(name, "won against", game.Dealer.Name)
			return
		}
	}
}
