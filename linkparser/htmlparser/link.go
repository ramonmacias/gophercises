package htmlparser

import (
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Link will keep all information related with (<a href="...."/>) a tag from
// HTML document
type Link struct {
	Href string
	Text string
}

// ParseFile from a filename we will parse into a list of Links
func ParseFile(filename string) ([]Link, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return Parse(f)
}

// ParseValue given a HTML doc in string format we will convert it into
// slice of Links
func ParseValue(htmlDoc string) ([]Link, error) {
	return Parse(strings.NewReader(htmlDoc))
}

// Parse given a Reader interface we are going to parse and get a slice of
// Links
func Parse(r io.Reader) (links []Link, err error) {
	doc, err := html.Parse(r)

	var funcNode func(*html.Node)
	var currentParent *html.Node

	funcNode = func(n *html.Node) {
		if (currentParent == n.Parent || n.Parent.Parent == currentParent) && len(links) > 0 {
			if n.Type == html.TextNode {
				links[len(links)-1].Text += n.Data
			}
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, Link{Href: a.Val})
					currentParent = n
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			funcNode(c)
		}
	}
	funcNode(doc)
	return links, nil
}
