package main

import (
	"fmt"
	"time"

	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/game"
	"github.com/RodrigoCelso/gophercises-10/player"
)

func main() {
	game := game.Game{Shoe: *deck.New(deck.WithMultipleDeckSize(4), deck.WithShuffle()), Dealer: player.New(), Players: []*player.Player{player.New()}}
	game.StartRound()
	fmt.Println("My first card")
	userPlay(&game, game.Players[0])
	dealerPlay(&game, game.Dealer)
}

func userPlay(game *game.Game, user *player.Player) {
	var userChoice int
	fmt.Println("What are you going to do?")
	fmt.Println("1. Hit")
	fmt.Println("2. Stand")
	fmt.Scanf("%d", &userChoice)
	fmt.Println(userChoice)
	switch userChoice {
	case 1:
		game.Hit(user)
	case 2:
		game.Stand(user)
	default:
		fmt.Println("Invalid action")
	}
	fmt.Println("Your Score:")
}

func dealerPlay(game *game.Game, dealer *player.Player) {
	fmt.Println("My second card:", dealer.Hand[1])
	score := dealer.Score()
	fmt.Println("Score: ", score)
	time.Sleep(2 * time.Second)
	for score <= 16 {
		newCard := game.Hit(dealer)
		score = dealer.Score()
		fmt.Println("Hit:", newCard, "Score:", score)
		time.Sleep(2 * time.Second)
	}
	fmt.Println("Stand:", score)
}
