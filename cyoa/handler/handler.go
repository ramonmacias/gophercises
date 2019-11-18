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

	chapter := story.GetChapter(mux.Vars(req)["id"])

	err = t.Execute(resp, chapter)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
