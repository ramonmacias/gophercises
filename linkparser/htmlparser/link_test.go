package htmlparser_test

import (
	"testing"

	"github.com/ramonmacias/gophercises/linkparser/htmlparser"
)

func TestParseFile(t *testing.T) {
	htmlparser.ParseFile("testdata/ex1.html")
}
