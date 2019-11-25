package main

import (
	"github.com/ramonmacias/gophercises/phone/db"
	"github.com/ramonmacias/gophercises/phone/domain"
)

func main() {
	db.Start()
	err := db.Clean()
	check(err)
	err = domain.CreatePhoneTable()
	check(err)
	db.Stop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
