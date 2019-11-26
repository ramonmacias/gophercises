package domain

import (
	"database/sql"
	"fmt"

	"github.com/ramonmacias/gophercises/phone/db"
)

// All the queries we are going to use for manage all related with the phones
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
	updatePhone   = `UPDATE phone SET original_number = '%s' , normalized_number = '%s' WHERE id = %d;`
)

// Phone defines the concept of phone using fields for save original and normalized numbers
type Phone struct {
	ID               int
	OriginalNumber   string
	NormalizedNumber string
}

// Save method will create a new phone into our database
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

// Update method will update all the values for the given phone
func (p *Phone) Update() error {
	tx, err := db.GetClient().BeginTx(db.GetContext(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	_, err = tx.Exec(fmt.Sprintf(updatePhone, p.OriginalNumber, p.NormalizedNumber, p.ID))
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

// ListAllPhones function will return the list of all the phones saved into our db
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

// CreatePhoneTable will create the table where we save phones
func CreatePhoneTable() error {
	_, err := db.GetClient().Query(createPhoneTable)
	if err != nil {
		return err
	}
	return nil
}
