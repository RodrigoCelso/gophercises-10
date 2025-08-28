package main

import (
	"fmt"

	"github.com/RodrigoCelso/gophercises-10/deck"
)

func main() {
	bwDeck := deck.New(
		deck.WithJoker(3),
		deck.WithShuffle(),
		deck.WithFilter(
			func(card deck.Card) bool {
				return card.CompareSuit("Joker") || card.CompareSuit("Club") || card.CompareValue("Two")
			}))
	bwDeck2 := deck.New(deck.WithMultipleDeck(*bwDeck))

	for _, card := range *bwDeck2 {
		// suit, deck := deck.ParseBitwiseCard(card)
		// fmt.Printf("%032b", card.CardValue)
		fmt.Println(card)
	}
}
