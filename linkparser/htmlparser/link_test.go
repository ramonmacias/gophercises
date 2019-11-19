package htmlparser_test

import (
	"testing"

	"github.com/ramonmacias/gophercises/linkparser/htmlparser"
)

func TestParseFile(t *testing.T) {
	expectedText := "A link to another page"
	expectedHref := "/other-page"

	links, err := htmlparser.ParseFile("testdata/ex1.html")
	checkLinksSizeAndError(1, links, err, t)
	checkExpectedLink(expectedText, expectedHref, links[0].Text, links[0].Href, t)
}

func TestParseSecondFile(t *testing.T) {
	links, err := htmlparser.ParseFile("testdata/ex2.html")
	checkLinksSizeAndError(2, links, err, t)
	checkExpectedLink("Check me out on twitter", "https://www.twitter.com/joncalhoun", links[0].Text, links[0].Href, t)
	checkExpectedLink("Gophercises is on Github", "https://github.com/gophercises", links[1].Text, links[1].Href, t)
}

func TestParseThirdFile(t *testing.T) {
	links, err := htmlparser.ParseFile("testdata/ex3.html")
	checkLinksSizeAndError(3, links, err, t)
}

func TestParseFileWithComment(t *testing.T) {
	expectedText := "dog cat "
	expectedHref := "/dog-cat"

	links, err := htmlparser.ParseFile("testdata/ex4.html")
	checkLinksSizeAndError(1, links, err, t)
	checkExpectedLink(expectedText, expectedHref, links[0].Text, links[0].Href, t)
}

func checkExpectedLink(expectedText, expectedHref, gotText, gotHref string, t *testing.T) {
	if gotHref != expectedHref {
		t.Errorf("Expected Href %s but got %s", expectedHref, gotHref)
	}
	if gotText != expectedText {
		t.Errorf("Expected Text as a %s but got %s", expectedText, gotText)
	}
}

func checkLinksSizeAndError(expectedSize int, links []htmlparser.Link, err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error should be nil but got %v", err)
	}
	if len(links) != expectedSize {
		t.Errorf("Length of links should be %d but got %d", expectedSize, len(links))
	}
}
