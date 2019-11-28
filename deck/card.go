package deck

import "sort"

//go:generate stringer -type=Suit,Rank

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var Suits = []Suit{Spade, Diamond, Club, Heart}

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

var Ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

type Card struct {
	Rank
	Suit
}

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

func Sort(cards []Card, sortFunc func(cards []Card) []Card) []Card {
	if sortFunc != nil {
		return sortFunc(cards)
	}
	return defaultSortFunc(cards)
}

func defaultSortFunc(cards []Card) []Card {
	sort.Slice(cards, func(i int, j int) bool {
		if i < j {
			return true
		}
		return false
	})
	return cards
}
