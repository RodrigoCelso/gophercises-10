package database

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

func ConnectDB() (*bolt.DB, error) {
	DB, err := bolt.Open("scores.db", 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("failed to access the database: %w", err)
	}
	return DB, nil
}
