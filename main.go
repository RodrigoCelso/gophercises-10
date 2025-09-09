package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/RodrigoCelso/gophercises-10/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/database"
	"github.com/RodrigoCelso/gophercises-10/deck"
	"github.com/RodrigoCelso/gophercises-10/game"
	"github.com/boltdb/bolt"
)

const FOCUS_TIME = 2 * time.Second

var playerMaxScore int

func main() {
	var round *game.Game
	playerMaxScore = 0

	for {
		fmt.Println("Blackjack\n0. Exit\n1. Start new Game?\n2. Scoreboard")
		var menuChoice uint8
		fmt.Scanf("%d\n", &menuChoice)
		if menuChoice == 0 {
			return
		} else if menuChoice == 1 {
			fmt.Println("How many players?")
			var playersQuantityChoice uint8
			fmt.Scanf("%d\n", &playersQuantityChoice)
			round = game.New(int(playersQuantityChoice))
			break
		} else if menuChoice == 2 {
			scoreboard()
		} else {
			continue
		}
	}

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
	dealCards(round.Players, round.Dealer, &round.Shoe)

	// Dealer reveals first card
	fmt.Println("this is my first card:")
	fmt.Print(dealerFlipCard(round.Dealer, 1), "\n\n")
	time.Sleep(FOCUS_TIME)

	// Players turn
	for _, p := range round.Players {
		playerPlay(round, p)
		time.Sleep(FOCUS_TIME)
	}

	// Dealer reveals second card
	fmt.Println("this is my second card:")
	fmt.Print(dealerFlipCard(round.Dealer, 2), "\n\n")
	fmt.Println("my current hand is:\n", round.Dealer)
	time.Sleep(FOCUS_TIME)

	// Dealer turn
	dealerPlay(round)
	time.Sleep(FOCUS_TIME)

	// Decide winners
	compareCards(round.Players, round.Dealer)
}

// OK
func dealCards(players []*bjackplayer.Player, dealer *bjackplayer.Player, shoe *deck.Deck) {
	for _, p := range players {
		p.Hit(shoe)
		p.Hit(shoe)
		pScore := p.Score()
		if pScore == 21 {
			// Natural Blackjack
			p.State = bjackplayer.Blackjack
		}
		if pScore > playerMaxScore {
			playerMaxScore = pScore
		}
	}
	dealer.Hit(shoe)
	dealer.Hit(shoe)
	if dealer.Score() == 21 {
		// Natural Blackjack
		dealer.State = bjackplayer.Blackjack
	}
}

// OK
func dealerFlipCard(dealer *bjackplayer.Player, cardNumber int) *deck.Card {
	return &dealer.Hand[cardNumber-1]
}

func playerPlay(round *game.Game, p *bjackplayer.Player) {
	fmt.Println("=====================\n", p, "\n=====================")
	for {
		fmt.Printf("Would you like to Hit or Stand? (Current Score: %d)\n", p.Score())
		fmt.Println("1. Hit\n2. Stand\n3. Check table")

		var playerChoice int
		fmt.Scanf("%d\n", &playerChoice)

		if playerChoice == 1 {
			cardHitted := p.Hit(&round.Shoe)
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
		} else if playerChoice == 3 {
			checkTable(round, p)
		} else {
			// invalid command
			fmt.Println("Invalid command, please try again.")
			continue
		}
	}
	fmt.Print("Final Score: ", p.Score(), "\n=====================\n")
}

func checkTable(round *game.Game, pCall *bjackplayer.Player) {
	fmt.Println("=====================")
	fmt.Println(round.Dealer)
	fmt.Println("=====================")
	fmt.Println("\n=====================")
	for _, p := range round.Players {
		fmt.Println(p)
		if p == pCall {
			fmt.Println("======= You ^ =======")
			continue
		}
		fmt.Println("=====================")
	}
}

func dealerPlay(round *game.Game) {
	score := round.Dealer.Score()
	for score < 17 || (score < playerMaxScore && len(round.Players) == 1) {
		cardHitted := round.Dealer.Hit(&round.Shoe)
		score = round.Dealer.Score()
		fmt.Print("\nI hitted: ", cardHitted, " - Value: ", cardHitted.BlackjackScore(), "\n")
		fmt.Print("My current score: ", score, "\n\n")
		time.Sleep(FOCUS_TIME)
		if score > playerMaxScore {
			break
		}
	}
	if score > 21 {
		fmt.Print("Dealer BUSTED ", round.Dealer.Score(), "\n")
		round.Dealer.Hand = []deck.Card{}
	}
}

func compareCards(players []*bjackplayer.Player, dealer *bjackplayer.Player) {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer db.Close()

	dealerScore := dealer.Score()
	for _, p := range players {
		pScore := p.Score()
		if (pScore == dealerScore) || (p.State == dealer.State && !(p.State == bjackplayer.Default && dealer.State == bjackplayer.Default)) {
			// tie
			fmt.Println(dealer.Name, "and", p.Name, "tied")
			db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("scoreboard"))
				if err != nil {
					return err
				}
				bucketKey := p.Name + " - " + time.Now().Format("02/01/2006 15:04:05")
				bucketValue := "Tie\n"
				bucket.Put([]byte(bucketKey), []byte(bucketValue))
				return nil
			})
			continue
		}

		if dealer.State == bjackplayer.Blackjack || dealerScore > pScore {
			// dealer won
			fmt.Println(dealer.Name, "won against", p.Name)
			err := db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("scoreboard"))
				if err != nil {
					return err
				}
				bucketKey := p.Name + " - " + time.Now().Format("02/01/2006 15:04:05")
				bucketValue := "Lose - " + "Chips: " + strconv.Itoa(p.Bet) + "\n"
				bucket.Put([]byte(bucketKey), []byte(bucketValue))
				return nil
			})
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			continue
		}

		// dealer lost
		err := db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte("scoreboard"))
			if err != nil {
				return err
			}
			bucketKey := p.Name + " - " + time.Now().Format("02/01/2006 15:04:05")
			bucketValue := "Win - " + "Chips: " + strconv.Itoa(p.Bet) + "\n"
			bucket.Put([]byte(bucketKey), []byte(bucketValue))
			return nil
		})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(p.Name, "won against", dealer.Name)
	}
}

func scoreboard() {
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("scoreboard"))
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("%s : %s", k, v)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
