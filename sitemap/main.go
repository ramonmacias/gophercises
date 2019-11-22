package main

import (
	"encoding/xml"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ramonmacias/gophercises/linkparser/htmlparser"
)

// SiteMap keep the information about all the links from a website
type SiteMap struct {
	Url []Url `xml:"url"`
}

// Urlkeep the information and structure for the definition of a site URL
type Url struct {
	Loc string `xml:"loc"`
}

var (
	website = flag.String("website", "", "Inform here the website you want to build a sitemap")
)

func main() {
	flag.Parse()
	if *website == "" {
		panic("The website flag is mandatory, use go run main.go -website=$website_val")
	}

	siteMap, err := buildSiteMap(getAllLinks(*website, 3))
	if err != nil {
		panic(err)
	}

	output, err := xml.MarshalIndent(siteMap, " ", "    ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(output)

}

// getLinks will take a website url as a parameter and return a list of links
// found on this page
func getLinks(website string) (links []htmlparser.Link, err error) {
	resp, err := http.Get(website)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	q, err := htmlparser.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	for _, link := range q {
		l, err := cleanLink(website, link)
		if err != nil {
			return nil, err
		}
		if l != nil {
			links = append(links, *l)
		}
	}

	return links, nil
}

// cleanLink given a root website and a link we clean and discard all the links
// that are not part of the same domain
func cleanLink(website string, link htmlparser.Link) (*htmlparser.Link, error) {
	rootUrl, err := url.Parse(website)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(link.Href)
	if err != nil {
		return nil, err
	}
	if u.Host == "" {
		return &htmlparser.Link{
			Href: website + link.Href,
			Text: link.Text,
		}, nil
	} else if u.Host == rootUrl.Host {
		return &htmlparser.Link{
			Href: link.Href,
			Text: link.Text,
		}, nil
	}
	return nil, nil
}

// getAllLinks will return all the links from a website url given as a parameter
// and all the childrens base on a depth level n
func getAllLinks(website string, n int) (linksRes []htmlparser.Link) {
	rootLinks, _ := getLinks(website)
	for _, link := range rootLinks {
		linksRes = append(linksRes, getChildLinksWithDepthLimit(link, n)...)
	}
	return linksRes
}

// getChildLinksWithDepthLimit given a root and a depth limit this function will
// find all the links, I used the BFS algorithm to achieve this
func getChildLinksWithDepthLimit(root htmlparser.Link, n int) (linksRes []htmlparser.Link) {
	discovered := make(map[string]bool)
	depthCount := 0
	q := []htmlparser.Link{}

	discovered[root.Href] = true
	q = append(q, root)
	log.Printf("Root Link: %s", root.Href)
	linksRes = append(linksRes, root)
	for len(q) > 0 && depthCount < n {
		v := q[0]
		q := q[1:]

		childs, _ := getLinks(v.Href)
		for _, link := range childs {
			if !discovered[link.Href] {
				log.Println(link.Href)
				linksRes = append(linksRes, link)
				discovered[link.Href] = true
				q = append(q, link)
			}
		}
		depthCount++
	}

	return linksRes
}

// buildSiteMap from a list of links will build a sitemap struct that then we can
// use it to marshal into a xml
func buildSiteMap(links []htmlparser.Link) (siteMap *SiteMap, err error) {
	siteMap = &SiteMap{}
	for _, link := range links {
		siteMap.Url = append(siteMap.Url, Url{link.Href})
	}
	return siteMap, err
}
