package domain

import (
	"database/sql"
	"fmt"

	"github.com/ramonmacias/gophercises/phone/db"
)

const (
	createPhoneTable = `
    CREATE TABLE phone(
     id serial PRIMARY KEY,
     original_number VARCHAR (50) NOT NULL,
     normalized_number VARCHAR (50) NOT NULL
     );`
	insertPhone = `
    INSERT INTO phone(original_number, normalized_number)
    VALUES ('%s', '%s');`
	listAllPhones = `select id, original_number, normalized_number from phone;`
)

type Phone struct {
	ID               int
	OriginalNumber   string
	NormalizedNumber string
}

func (p *Phone) Save() error {
	tx, err := db.GetClient().BeginTx(db.GetContext(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	_, err = tx.Exec(fmt.Sprintf(insertPhone, p.OriginalNumber, p.NormalizedNumber))
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func ListAllPhones() (phones []Phone, err error) {
	rows, err := db.GetClient().QueryContext(db.GetContext(), listAllPhones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		phone := Phone{}
		if err = rows.Scan(&phone.ID, &phone.OriginalNumber, &phone.NormalizedNumber); err != nil {
			return nil, err
		}
		phones = append(phones, phone)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return phones, nil
}

func CreatePhoneTable() error {
	_, err := db.GetClient().Query(createPhoneTable)
	if err != nil {
		return err
	}
	return nil
}
