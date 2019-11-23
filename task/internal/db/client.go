package db

import (
	"sync"

	"github.com/boltdb/bolt"
)

var (
	c    *configuration
	once sync.Once
)

type configuration struct {
	client *bolt.DB
}

func Start() {
	setup()
}

func setup() *configuration {
	once.Do(func() {
		c = &configuration{
			client: boltClient(),
		}
	})
	return c
}

func boltClient() (client *bolt.DB) {
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	return db
}

func GetClient() *bolt.DB {
	return setup().client
}
