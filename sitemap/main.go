package main

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"

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

	getAllLinks(html, 0)

	// links, err := htmlparser.Parse(html)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// siteMap, err := buildSiteMap(links)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// output, err := xml.MarshalIndent(siteMap, " ", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// os.Stdout.Write(output)
}

func getHtmlPage(website *string) (r io.Reader, err error) {
	resp, err := http.Get(*website)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func getAllLinks(r io.Reader, n int) {
	links, err := htmlparser.Parse(r)
	if err != nil {
		panic(err)
	}
	if n >= 3 {
		log.Println("BREAAAK")
		return
	} else {
		for _, link := range links {
			log.Println(link.Href)
			f, _ := getHtmlPage(&link.Href)
			if f != nil {
				n++
				getAllLinks(f, n)
			}
		}
	}
}

func buildSiteMap(links []htmlparser.Link) (siteMap *SiteMap, err error) {
	siteMap = &SiteMap{}
	for _, link := range links {
		siteMap.Url = append(siteMap.Url, Url{link.Href})
		html, err := getHtmlPage(&link.Href)
		if err == nil {
			links2, err := htmlparser.Parse(html)
			if err != nil {
				panic(err)
			}
			for _, link2 := range links2 {
				siteMap.Url = append(siteMap.Url, Url{link2.Href})
			}
		}
	}
	return siteMap, err
}
