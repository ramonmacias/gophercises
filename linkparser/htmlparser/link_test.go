package htmlparser_test

import (
	"testing"

	"github.com/ramonmacias/gophercises/linkparser/htmlparser"
)

func TestParseFile(t *testing.T) {
	expectedText := "A link to another page"
	expectedHref := "/other-page"

	links, err := htmlparser.ParseFile("testdata/ex1.html")
	if err != nil {
		t.Errorf("Error should be nil but got %v", err)
	}
	if len(links) != 1 {
		t.Errorf("Length of links should be 1 but got %d", len(links))
	}
	if links[0].Href != expectedHref {
		t.Errorf("Expected Href %s but got %s", expectedHref, links[0].Href)
	}
	if links[0].Text != expectedText {
		t.Errorf("Expected Text as a %s but got %s", expectedText, links[0].Text)
	}
}
