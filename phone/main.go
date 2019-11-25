package main

import (
	"log"

	"github.com/ramonmacias/gophercises/phone/db"
	"github.com/ramonmacias/gophercises/phone/domain"
)

func main() {
	db.Start()
	err := db.Clean()
	check(err)
	err = domain.CreatePhoneTable()
	check(err)

	phone := domain.Phone{
		OriginalNumber:   "444455(55)",
		NormalizedNumber: "44445555",
	}
	err = phone.Save()
	check(err)
	secondPhone := domain.Phone{
		OriginalNumber:   "11111(55)",
		NormalizedNumber: "11111222",
	}
	err = secondPhone.Save()
	check(err)

	phones, err := domain.ListAllPhones()
	check(err)
	for _, phone := range phones {
		log.Printf("Phone ID: %d original number: %s normalized number: %s\n", phone.ID, phone.OriginalNumber, phone.NormalizedNumber)
	}
	db.Stop()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
