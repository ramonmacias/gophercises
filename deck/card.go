package deck

import (
	"math/rand"
	"sort"
	"time"
)

//go:generate stringer -type=Suit,Rank

const (
	maxNumberJokers = 10
)

// Suit type represents the which can of suit has the card
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

// Suits is an array that keeps all the existing suits
var Suits = []Suit{Spade, Diamond, Club, Heart}

// Rank type will represent the value for each card
type Rank uint8

const (
	_ Rank = iota
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
)

// Ranks keep the value of all the Ranks created
var Ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Card type represent the information of a Card
type Card struct {
	Rank
	Suit
}

// New function will build a deck, that is an slice of cards with a size of len(ranks)
// * len(suits)
func New() (cards []Card) {
	for _, suit := range Suits {
		for _, rank := range Ranks {
			cards = append(cards, Card{
				Suit: suit,
				Rank: rank,
			})
		}
	}
	return cards
}

// Sort function will sort an slice of cards with a given function sort or use
// a default function instead
func Sort(cards []Card, sortFunc func(cards []Card) []Card) []Card {
	if sortFunc != nil {
		return sortFunc(cards)
	}
	return defaultSortFunc(cards)
}

// defaultSortFunc is the default sorting function, on that case will sort the desk
// from less to max in base of the rank value
func defaultSortFunc(cards []Card) []Card {
	sort.Slice(cards, func(i int, j int) bool {
		if cards[i].Rank < cards[j].Rank {
			return true
		}
		return false
	})
	return cards
}

// Shuffle functions just re sort the slice into a random order, using the current
// time in nanosecond as a seed
func Shuffle(cards []Card) []Card {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	r.Shuffle(len(cards), func(i int, j int) {
		aux := cards[i]
		cards[i] = cards[j]
		cards[j] = aux
	})
	return cards
}

// AddRandomJokers is a function that will add new joker cards into a given deck
// from a number between 0 and maxNumberJokers constant
func AddRandomJokers(cards []Card) []Card {
	return AddRandomJokersN(cards, maxNumberJokers)
}

// AddRandomJokersN is a function that will add a maximum of n random jokers
func AddRandomJokersN(cards []Card, n int) []Card {
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for index := 0; index < r.Intn(n)+1; index++ {
		cards = append(cards, Card{Ace, Joker})
	}
	return cards
}
