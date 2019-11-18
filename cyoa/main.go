package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramonmacias/gophercises/cyoa/handler"
	"github.com/ramonmacias/gophercises/cyoa/story"
)

var (
	mode = flag.String("mode", "server", "Select if we want to see the story in html or in the terminal")
)

func main() {
	flag.Parse()
	switch *mode {
	case "server":
		runServer()
	case "terminal":
		runTerminal()
	default:
		panic("Mode not allowed")
	}
}

func runServer() {
	m := mux.NewRouter()
	m.HandleFunc("/chapter/{id}", handler.ChapterHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", m)
}

func runTerminal() {
	chapter := story.GetChapter("intro")
	log.Println("Welcome to the Gopher choose your own adventure!")

	for len(chapter.Options) > 0 {
		for _, paragraph := range chapter.Paragraphs {
			log.Println(paragraph)
		}
		log.Println("Please choose one of the next options")
		for i, option := range chapter.Options {
			log.Printf("%d - %s", i+1, option.Text)
		}
		var optionSelected int
		_, err := fmt.Scanln(&optionSelected)
		if err != nil {
			panic(err)
		}
		chapter = story.GetChapter(chapter.Options[optionSelected-1].Arc)
	}

	for _, paragraph := range chapter.Paragraphs {
		log.Println(paragraph)
	}

	log.Println("THANKS FOR READING OUR ADVENTURE!")
}
