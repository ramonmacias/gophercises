package normalizer_test

import (
	"testing"

	"github.com/ramonmacias/gophercises/phone/normalizer"
)

func TestFilterParenthesis(t *testing.T) {
	want := "122233"
	input := "1222(33)"

	got := normalizer.FilterParenthesis()(input)
	if want != got {
		t.Errorf("We want %s but got %s", want, got)
	}
}

func TestFilterHyphens(t *testing.T) {
	want := "122233"
	input := "12-223-3"

	got := normalizer.FilterHyphens()(input)
	if want != got {
		t.Errorf("We want %s but got %s", want, got)
	}
}

func TestFilterFuncs(t *testing.T) {
	want := "11222333"
	phoneNumber := "11-22-2(333)"

	filters := normalizer.FilterFuncs()
	for _, filter := range filters {
		phoneNumber = filter(phoneNumber)
	}
	if want != phoneNumber {
		t.Errorf("We want %s but got %s", want, phoneNumber)
	}
}
