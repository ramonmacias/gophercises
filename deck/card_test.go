package deck_test

import (
	"log"
	"testing"

	"github.com/ramonmacias/gophercises/deck"
)

func TestBuilder(t *testing.T) {
	c := deck.Card{
		Suit: deck.Spade,
		Rank: deck.Ace,
	}

	log.Println(c, c.Suit, c.Rank)
}

func TestFuncNew(t *testing.T) {
	expectedLength := 52
	d := deck.New()
	if len(d) != expectedLength {
		t.Errorf("Expected %d but got %d", expectedLength, len(d))
	}
}

func TestSort(t *testing.T) {
	d := deck.New()
	log.Println(d)
	deck.Sort(d, nil)
	log.Println(d)
}
