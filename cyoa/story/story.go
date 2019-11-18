package story

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type chapter struct {
	Title      string
	Paragraphs []string `json:"story"`
	Options    []option
}

type option struct {
	Text string
	Arc  string
}

var (
	story map[string]chapter
)

func init() {
	f, err := os.Open("story/gopher.json")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(b, &story); err != nil {
		panic(err)
	}
}

func GetChapter(id string) chapter {
	return story[id]
}
