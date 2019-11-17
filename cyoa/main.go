package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	tpl = `
					<!DOCTYPE html>
					<html>
						<head>
							<meta charset="UTF-8">
							<title>{{.Title}}</title>
						</head>
						<body>
							{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
						</body>
					</html>`
)

func main() {
	m := defaultMux()
	m.HandleFunc("/init", func(resp http.ResponseWriter, req *http.Request) {
		t, err := template.New("webpage").Parse(tpl)
		check(err)

		data := struct {
			Title string
			Items []string
		}{
			Title: "My page",
			Items: []string{
				"My photos",
				"My blog",
			},
		}

		err = t.Execute(resp, data)
		check(err)

		noItems := struct {
			Title string
			Items []string
		}{
			Title: "My another page",
			Items: []string{},
		}

		err = t.Execute(resp, noItems)
		check(err)
	})

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", m)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
