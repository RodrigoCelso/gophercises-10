package controller

import (
	"fmt"
	"strconv"

	"github.com/RodrigoCelso/gophercises-10/internal/database"
	"github.com/boltdb/bolt"
)

func GetChips(playerName string) (int, error) {
	pChips := 0

	db, err := database.ConnectDB()
	if err != nil {
		return 0, fmt.Errorf("a problem occured while oppening the database: %w", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("chips"))
		if err != nil {
			return err
		}

		chips := bucket.Get([]byte(playerName))
		if chips != nil {
			pChips, err = strconv.Atoi(string(chips))
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("a problem occured while reading the database: %w", err)
	}

	return pChips, nil
}

func InsertRemoveChips(playerName string, chips int) error {
	db, err := database.ConnectDB()
	if err != nil {
		return fmt.Errorf("a problem occured while oppening the database: %w", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("chips"))
		if err != nil {
			return err
		}

		playerChipsStr := string(bucket.Get([]byte(playerName)))
		var playerChips int

		if playerChipsStr != "" {
			playerChips, err = strconv.Atoi(playerChipsStr)
			if err != nil {
				return err
			}
		}

		err = bucket.Put([]byte(playerName), []byte(strconv.Itoa(playerChips+chips)))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("a problem occured while inserting to database: %w", err)
	}
	return nil
}
