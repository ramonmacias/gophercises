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
	// Flag used for select between a server or terminal mode
	mode = flag.String("mode", "server", "Select if we want to see the story in html (server) or in the terminal (terminal)")
)

// Function main from where we run the program in a server or in a terminal
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

// runServer function will build a new server with the handler for show using go
// templates the story in an html page
func runServer() {
	m := mux.NewRouter()
	m.HandleFunc("/chapter/{id}", handler.ChapterHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", m)
}

// runTerminal will show the same story but from the terminal instead from an
// html file
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
