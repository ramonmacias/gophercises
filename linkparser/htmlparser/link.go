package htmlparser

import (
	"log"
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
	funcNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					log.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			funcNode(c)
		}
	}
	funcNode(doc)
	return nil, nil
}
