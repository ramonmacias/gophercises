package normalizer_test

import (
	"reflect"
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

func TestBatchNormalize(t *testing.T) {
	want := []string{"333444", "55566666", "22122332"}
	input := []string{"3334-4-4", "5(556)66-66", "(221)22-33-2"}

	got := normalizer.BatchNormalize(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("We want %s but got %s", want, got)
	}
}
