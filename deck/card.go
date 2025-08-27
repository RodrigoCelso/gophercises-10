//go:generate stringer -type=CardSuit,CardValue
package deck

import (
	"math/rand"
	"sort"
)

// CardSuit represents the suit from the Card (Spade = 0, ...).
// Joker is considered a suit.
type CardSuit int

const (
	Spade CardSuit = iota
	Diamond
	Club
	Heart
	suitSize
	Joker
)

// CardValue are the value written in the card (Ace=1, ...)
type CardValue int

const (
	_ CardValue = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	valueSize
)

// Cards are the structures that composes the deck.
type Card struct {
	Suit       CardSuit
	Value      CardValue
	TrumpValue int
}

func newSuitDeck(suit CardSuit) []Card {
	var suitDeck []Card
	for value := range valueSize - 1 {
		card := Card{Suit: suit, Value: CardValue(value + 1)}
		suitDeck = append(suitDeck, card)
	}
	return suitDeck
}

// New allocates a new slice of Card given the n Options.
func New(opts ...option) *[]Card {
	var deck []Card
	for suit := range suitSize {
		deck = append(deck, newSuitDeck(CardSuit(suit))...)
	}
	for _, o := range opts {
		deck = o(deck)
	}
	return &deck
}

type option func([]Card) []Card

// WithJoker appends the given quantity of Joker cards.
func WithJoker(quantity int) option {
	return func(c []Card) []Card {
		for idx := 0; idx < quantity; idx++ {
			card := Card{Suit: Joker, Value: CardValue(0), TrumpValue: 55}
			c = append(c, card)
		}
		return c
	}
}

// WithSort returns the cards sorted by suit.
func WithSort() option {
	return func(c []Card) []Card {
		sort.Slice(c, func(i, j int) bool {
			return c[i].Suit < c[j].Suit
		})
		return c
	}
}

// WithSortFunc returns the cards sorted by the given function.
func WithSortFunc(less func(cards []Card) func(i, j int) bool) option {
	return func(c []Card) []Card {
		sort.Slice(c, less(c))
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
