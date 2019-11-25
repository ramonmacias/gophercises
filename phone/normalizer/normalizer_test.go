package normalizer_test

import (
	"testing"

	"github.com/ramonmacias/gophercises/phone/normalizer"
)

func TestNormalize(t *testing.T) {
	want := "111222333"
	input := "111-2(22-33(3)"

	got := normalizer.Normalize(input)
	if want != got {
		t.Errorf("We want %s but got %s", want, got)
	}
}
