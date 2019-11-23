package caesarcipher_test

import (
	"testing"

	"github.com/ramonmacias/gophercises/hackerrank/caesarcipher"
)

func TestEncrypt(t *testing.T) {
	want := "okffng-Qwvb"
	input := "middle-Outz"
	got := caesarcipher.Encrypt(input, 2)
	if want != got {
		t.Errorf("We want %s but got %s", want, got)
	}

	// want = "fff.jkl.gh"
	// input = "www.abc.xy"
	// got = caesarcipher.Encrypt(input, 10)
	// if want != got {
	// 	t.Errorf("We want %s but got %s", want, got)
	// }
}

func TestEncryptWithASCII(t *testing.T) {
	want := "okffng-Qwvb"
	input := "middle-Outz"
	got := caesarcipher.EncryptASCII(input, 2)
	if want != got {
		t.Errorf("We want %s but got %s", want, got)
	}

	want = "fff.jkl.gh"
	input = "www.abc.xy"
	got = caesarcipher.EncryptASCII(input, 10)
	if want != got {
		t.Errorf("We want %s but got %s", want, got)
	}
}
