package deck

import (
	"fmt"
	"math/bits"
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
	Suit  uint8
	Value uint16
}

type Deck []Card

func (d Deck) String() string {
	var cards string
	for _, card := range d {
		cards += fmt.Sprintln(card)
	}
	return cards
}

type option func(Deck) Deck

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
func New(opts ...option) *Deck {
	var deck Deck
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
	return func(cb Deck) Deck {
		sort.Slice(cb, func(i, j int) bool {
			return cb[i].Suit < cb[j].Suit
		})
		return cb
	}
}

// WithSortFunc returns the cards sorted by the given function.
func WithSortFunc(less func(cards Deck) func(i, j int) bool) option {
	return func(cb Deck) Deck {
		sort.Slice(cb, less(cb))
		return cb
	}
}

// WithJoker appends the given quantity of Joker cards.
func WithJoker(quantity int) option {
	return func(c Deck) Deck {
		for idx := 0; idx < quantity; idx++ {
			card := Card{Suit: uint8(1 << 4), Value: uint16(0)}
			c = append(c, card)
		}
		return c
	}
}

// WithShuffle uses the rand standard package and returns the shuffled cards.
func WithShuffle() option {
	return func(c Deck) Deck {
		rand.Shuffle(len(c), func(i, j int) {
			c[i], c[j] = c[j], c[i]
		})
		return c
	}
}

// WithFilter returns a new filtered deck, removing all the cards that returns true by the function cardFilter.
func WithFilter(cardFilter func(card Card) bool) option {
	return func(cards Deck) Deck {
		var filteredDeck Deck
		for _, card := range cards {
			if !cardFilter(card) {
				filteredDeck = append(filteredDeck, card)
			}
		}
		return filteredDeck
	}
}

// WithMultipleDeck returns a new deck composed by the given decks.
func WithMultipleDeck(decks ...Deck) option {
	return func(c Deck) Deck {
		var multipleDeck Deck
		for _, deck := range decks {
			multipleDeck = append(multipleDeck, deck...)
		}
		return multipleDeck
	}
}

func WithMultipleDeckSize(size int) option {
	return func(c Deck) Deck {
		var multipleDeck Deck
		for range size {
			multipleDeck = append(multipleDeck, *New()...)
		}
		return multipleDeck
	}
}

func (d *Deck) PopCard() Card {
	cardsLen := len(*d)
	cardOut := (*d)[cardsLen-1]
	*d = (*d)[:cardsLen-1]
	return cardOut
}

func (d Deck) Shuffle() Deck {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
	return d
}

func (c *Card) BlackjackScore() int {
	score, _ := c.BlackjackScoreWithAce()
	return score
}

func (d *Deck) BlackjackScore() int {
	var aces uint8
	var scoreValue int
	for _, card := range *d {
		cardValue, isAce := card.BlackjackScoreWithAce()
		scoreValue += cardValue
		if isAce {
			aces++
		}
	}
	for range aces {
		if scoreValue > 21 {
			scoreValue -= 10
			aces--
		} else {
			aces--
		}
	}
	return scoreValue
}

func (c *Card) BlackjackScoreWithAce() (int, bool) {
	if c.Value == 0 {
		return 0, false
	}

	var isAce bool
	var scoreValue int
	cardValue := bits.TrailingZeros16(c.Value) + 1
	switch cardValue {
	case 1:
		isAce = true
		scoreValue += 11
	case 11, 12, 13:
		scoreValue += 10
	default:
		scoreValue += cardValue
	}
	return scoreValue, isAce
}
