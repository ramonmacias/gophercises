package deck_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/ramonmacias/gophercises/deck"
)

func TesBuilder(t *testing.T) {
	expectedRank := "Ace"
	expectedSuit := "Spade"

	c := deck.Card{deck.Ace, deck.Spade}
	if c.Rank.String() != expectedRank {
		t.Errorf("Expected %s but got %s", expectedRank, c.Rank.String())
	}
	if c.Suit.String() != expectedSuit {
		t.Errorf("Expected %s but got %s", expectedSuit, c.Suit.String())
	}
}

func TestFuncNew(t *testing.T) {
	expectedLength := 52
	d := deck.New()
	if len(d) != expectedLength {
		t.Errorf("Expected %d but got %d", expectedLength, len(d))
	}
}

func TestDefaultSortFunc(t *testing.T) {
	d := []deck.Card{deck.Card{deck.King, deck.Spade}, deck.Card{deck.Ace, deck.Spade}, deck.Card{deck.Two, deck.Spade}}
	expected := []deck.Card{deck.Card{deck.Ace, deck.Spade}, deck.Card{deck.Two, deck.Spade}, deck.Card{deck.King, deck.Spade}}

	d = deck.Sort(d, nil)
	if !reflect.DeepEqual(d, expected) {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func TestCustomSortFunc(t *testing.T) {
	d := []deck.Card{deck.Card{deck.King, deck.Spade}, deck.Card{deck.Ace, deck.Spade}, deck.Card{deck.Two, deck.Spade}}
	expected := []deck.Card{deck.Card{deck.King, deck.Spade}, deck.Card{deck.Two, deck.Spade}, deck.Card{deck.Ace, deck.Spade}}

	customSortFunc := func(cards []deck.Card) []deck.Card {
		sort.Slice(cards, func(i int, j int) bool {
			if cards[i].Rank > cards[j].Rank {
				return true
			}
			return false
		})
		return cards
	}

	d = deck.Sort(d, customSortFunc)

	if !reflect.DeepEqual(d, expected) {
		t.Errorf("Expected %v but got %v", expected, d)
	}
}

func TestShuffleFunc(t *testing.T) {
	d := deck.New()
	oldDeck := make([]deck.Card, len(d))
	copy(oldDeck, d)
	shuffled := deck.Shuffle(d)
	if reflect.DeepEqual(shuffled, oldDeck) {
		t.Error("After use shuffle function shouldn't have the same sorted slice")
	}
}

func TestAddRandomJokersFunc(t *testing.T) {
	d := deck.New()
	got := deck.AddRandomJokers(d)
	if len(got) == len(d) {
		t.Error("Can not have the same len if we add some jokers")
	}
}

func TestFilterFunc(t *testing.T) {
	d := deck.New()
	got := deck.Filter(d, []deck.Card{deck.Card{deck.Ace, deck.Spade}})
	if len(d) == len(got) {
		t.Errorf("Shouldn't be the same size, but expected %d and got %d", len(d)-1, len(got))
	}
	for _, card := range got {
		if card.Rank == deck.Ace && card.Suit == deck.Spade {
			t.Error("This card should be removed but isn't")
		}
	}
}

func TestNewMultiple(t *testing.T) {
	expectedSize := 3
	expectedInnerDeckSize := 52
	d := deck.NewMultiple(expectedSize)
	if len(d) != expectedSize {
		t.Errorf("Expected len %d but got %d", expectedSize, len(d))
	}
	for _, innerDeck := range d {
		if len(innerDeck) != expectedInnerDeckSize {
			t.Errorf("Expected len %d but got %d", expectedInnerDeckSize, len(innerDeck))
		}
	}
}
