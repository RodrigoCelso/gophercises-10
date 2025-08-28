package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Describes the amount of cards inside the deck (CARD_VALUES_SIZE * CARD_SUITS_SIZE = Deck Size).
const (
	CARD_VALUES_SIZE = 13
	CARD_SUITS_SIZE  = 4
)

// Cards are the structures that composes the deck.
type Card struct {
	Suit       uint8
	Value      uint16
	TrumpValue uint8
}

type option func([]Card) []Card

var cardSuitMap = map[uint8]string{
	uint8(1 << 0): "Spade",
	uint8(1 << 1): "Diamond",
	uint8(1 << 2): "Club",
	uint8(1 << 3): "Heart",
	uint8(1 << 4): "Joker",
}

var cardValueMap = map[uint16]string{
	uint16(1 << 0):  "Ace",
	uint16(1 << 1):  "Two",
	uint16(1 << 2):  "Three",
	uint16(1 << 3):  "Four",
	uint16(1 << 4):  "Five",
	uint16(1 << 5):  "Six",
	uint16(1 << 6):  "Seven",
	uint16(1 << 7):  "Eight",
	uint16(1 << 8):  "Nine",
	uint16(1 << 9):  "Ten",
	uint16(1 << 10): "Jack",
	uint16(1 << 11): "Queen",
	uint16(1 << 12): "King",
}

// Compare the card suit by the given suit string
func (c Card) CompareSuit(suit string) bool {
	s, ok := cardSuitMap[c.Suit]
	if !ok {
		return false
	}
	suit = strings.TrimSpace(suit)
	return strings.EqualFold(s, suit)
}

// Compare the card value by the given value string
func (c Card) CompareValue(value string) bool {
	v, ok := cardValueMap[c.Value]
	if !ok {
		return false
	}
	value = strings.TrimSpace(value)
	return strings.EqualFold(v, value)
}

// Returns the suit and value strings.
func TranscribeCard(card Card) (string, string) {
	return cardSuitMap[card.Suit], cardValueMap[card.Value]
}

// Prints the card's name based on how should spell it.
func (c Card) String() string {
	suit := cardSuitMap[c.Suit]
	value := cardValueMap[c.Value]
	if suit == "Joker" {
		return "Joker"
	}
	return fmt.Sprintf("%s of %ss", value, suit)
}

// New allocates a new slice of Card given the n Options.
func New(opts ...option) *[]Card {
	var deck []Card
	for suit := range CARD_SUITS_SIZE {
		for value := range CARD_VALUES_SIZE {
			newCardBW := Card{Value: 1 << value, Suit: 1 << suit}
			deck = append(deck, newCardBW)
		}
	}
	for _, opt := range opts {
		deck = opt(deck)
	}
	return &deck
}

// WithSort returns the cards sorted by suit.
func WithSort() option {
	return func(cb []Card) []Card {
		sort.Slice(cb, func(i, j int) bool {
			return cb[i].Suit < cb[j].Suit
		})
		return cb
	}
}

// WithSortFunc returns the cards sorted by the given function.
func WithSortFunc(less func(cards []Card) func(i, j int) bool) option {
	return func(cb []Card) []Card {
		sort.Slice(cb, less(cb))
		return cb
	}
}

// WithJoker appends the given quantity of Joker cards.
func WithJoker(quantity int) option {
	return func(c []Card) []Card {
		for idx := 0; idx < quantity; idx++ {
			card := Card{Suit: uint8(1 << 4), Value: uint16(0)}
			c = append(c, card)
		}
		return c
	}
}

// WithShuffle uses the rand standard package and returns the shuffled cards.
func WithShuffle() option {
	return func(c []Card) []Card {
		rand.Shuffle(len(c), func(i, j int) {
			c[i], c[j] = c[j], c[i]
		})
		return c
	}
}

// WithFilter returns the cards without the filtered cards by the given function.
func WithFilter(cardFilter func(card Card) bool) option {
	return func(cards []Card) []Card {
		var filteredDeck []Card
		for _, card := range cards {
			if !cardFilter(card) {
				filteredDeck = append(filteredDeck, card)
			}
		}
		return filteredDeck
	}
}

// WithMultipleDeck returns a new deck composed by the given decks.
func WithMultipleDeck(decks ...[]Card) option {
	return func(c []Card) []Card {
		var multipleDeck []Card
		for _, deck := range decks {
			multipleDeck = append(multipleDeck, deck...)
		}
		return multipleDeck
	}
}
