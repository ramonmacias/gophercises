package db

import (
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	c    *configuration
	once sync.Once
)

type configuration struct {
	ctx    context.Context
	client *sql.DB
	cancel context.CancelFunc
}

func Start() {
	setup()
}

func setup() *configuration {
	once.Do(func() {
		c = client()
	})
	return c
}

func client() *configuration {
	// TODO this should be moved into a .env file in order to allow a different connection
	connStr := "user=ramon dbname=phone_normalizer password=ramon_postgres_pass sslmode=disable"

	pool, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicf("Unable to data source name, cause %v", err)
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())

	if err = pool.PingContext(ctx); err != nil {
		log.Panicf("The is an error while try to do a ping to our database, cause %v", err)
	}

	return &configuration{
		ctx:    ctx,
		client: pool,
		cancel: stop,
	}
}

// Ping will send a ping to try the connection with the database
func Ping() error {
	return setup().client.PingContext(setup().ctx)
}

// GetClient will return the sql client
func GetClient() *sql.DB {
	return setup().client
}

// GetContext will return the context of db
func GetContext() context.Context {
	return setup().ctx
}

// Stop will close all the connections to our database
func Stop() {
	setup().cancel()
	setup().client.Close()
}

// Clean function will drop the phone table if exists
func Clean() error {
	tx, err := GetClient().BeginTx(GetContext(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	_, err = tx.Exec(`drop table if exists phone`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
