package gameflow

import (
	"fmt"
	"strconv"
	"time"

	"github.com/RodrigoCelso/gophercises-10/internal/bjackplayer"
	"github.com/RodrigoCelso/gophercises-10/internal/database"
	"github.com/boltdb/bolt"
)

func CompareCards(players []*bjackplayer.Player, dealer *bjackplayer.Player) {
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

func Scoreboard() {
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
