package domain

import "github.com/ramonmacias/gophercises/phone/db"

const (
	createPhoneTable = `
    CREATE TABLE phone(
     id serial PRIMARY KEY,
     original_number VARCHAR (50) NOT NULL,
     normalized_number VARCHAR (50) NOT NULL
     );`
)

type Phone struct {
	ID               int
	OriginalNumber   string
	NormalizedNumber string
}

func CreatePhoneTable() error {
	_, err := db.GetClient().Query(createPhoneTable)
	if err != nil {
		return err
	}
	return nil
}
