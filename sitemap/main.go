package main

import (
	"encoding/xml"
	"flag"
	"io"
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

	html, err := getHtmlPage(website)
	if err != nil {
		panic(err)
	}
	// TODO should be close due is a Body response
	// defer html.Close()

	links, err := htmlparser.Parse(html)
	if err != nil {
		panic(err)
	}

	siteMap, err := buildSiteMap(links)
	if err != nil {
		panic(err)
	}

	output, err := xml.MarshalIndent(siteMap, " ", "    ")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(output)
}

func getHtmlPage(website *string) (r io.Reader, err error) {
	resp, err := http.Get(*website)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func buildSiteMap(links []htmlparser.Link) (siteMap *SiteMap, err error) {
	siteMap = &SiteMap{}
	for _, link := range links {
		siteMap.Url = append(siteMap.Url, Url{link.Href})
	}
	return siteMap, err
}
