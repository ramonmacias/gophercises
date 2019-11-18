package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ramonmacias/gophercises/cyoa/story"
)

func ChapterHandler(resp http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("templates/chapter.html")
	check(err)

	data := struct {
		Title     string
		Paragraph string
		Items     []string
	}{
		Title:     mux.Vars(req)["id"],
		Paragraph: story.GetChapter(mux.Vars(req)["id"]).Title,
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(resp, data)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
