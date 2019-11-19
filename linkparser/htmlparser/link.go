package htmlparser

import (
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseFile(filename string) (links []Link, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(f)

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
