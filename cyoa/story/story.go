package story

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// chapter struct used for unmarshal json, with information about title and Paragraphs
type chapter struct {
	Title      string
	Paragraphs []string `json:"story"`
	Options    []option
}

// option struct used for unmarshal json with information detailed about options to choose
type option struct {
	Text string
	Arc  string
}

var (
	story map[string]chapter
)

// On the init function we unmarshal the json with all the story into our chapter
// struct and save it in memory
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

// GetChapter will return information about a specific chapter
func GetChapter(id string) chapter {
	return story[id]
}
