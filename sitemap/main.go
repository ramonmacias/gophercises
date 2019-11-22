package main

import (
	"encoding/xml"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ramonmacias/gophercises/linkparser/htmlparser"
)

type SiteMap struct {
	Url []Url `xml:"url"`
}

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

func getLinks(website string) ([]htmlparser.Link, error) {
	resp, err := http.Get(website)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	q, err := htmlparser.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return q, nil
}

func getAllLinks(website string, n int) (linksRes []htmlparser.Link) {
	rootLinks, _ := getLinks(website)
	log.Printf("Are going to search from %d root nodes", len(rootLinks))
	for _, link := range rootLinks {
		linksRes = append(linksRes, getChildLinksWithDepthLimit(link, n)...)
	}
	return linksRes
}

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

func buildSiteMap(links []htmlparser.Link) (siteMap *SiteMap, err error) {
	siteMap = &SiteMap{}
	for _, link := range links {
		siteMap.Url = append(siteMap.Url, Url{link.Href})
	}
	return siteMap, err
}
