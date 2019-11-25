package main

import "github.com/ramonmacias/gophercises/phone/db"

func main() {
	db.Start()
	db.Stop()
}
