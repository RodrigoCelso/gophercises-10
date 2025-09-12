package controller

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/internal/database"
	"github.com/boltdb/bolt"
)

func NewScoreboardEntry(key string, value string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return fmt.Errorf("a problem occured while oppening the database: %w", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("scoreboard"))
		if err != nil {
			return err
		}
		bucket.Put([]byte(key), []byte(value))
		return nil
	})
	if err != nil {
		return fmt.Errorf("a problem occured while updating the database: %w", err)
	}
	return nil
}

func Scoreboard() (map[string]string, error) {
	result := make(map[string]string)

	db, err := database.ConnectDB()
	if err != nil {
		return nil, fmt.Errorf("a problem occured while oppening the database: %w", err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("scoreboard"))
		if err != nil {
			return err
		}

		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			result[string(k)] = string(v)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("a problem occured while reading the database: %w", err)
	}
	return result, nil
}
