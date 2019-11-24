package db

import (
	"sync"

	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
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
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	db, err := bolt.Open(dir+"/task.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	return db
}

func GetClient() *bolt.DB {
	return setup().client
}
