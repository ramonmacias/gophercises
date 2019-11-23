package camelcase_test

import (
	"testing"

	"github.com/ramonmacias/gophercises/hackerrank/camelcase"
)

func TestCamelCaseWordCounter(t *testing.T) {
	want := int32(5)
	input := "saveChangesInTheEditor"

	got := camelcase.CamelCaseWordCounter(input)
	if want != got {
		t.Errorf("We want %d but got %d", want, got)
	}
}
