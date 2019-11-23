package domain

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
	"github.com/ramonmacias/gophercises/task/internal/db"
)

const (
	bucketName = "tasks"
)

// Task struct will keep information related with task
type Task struct {
	ID          int
	Description string
}

// Create method will create a new key value pair into our db
func (t *Task) Create() error {
	return db.GetClient().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		id, _ := b.NextSequence()
		t.ID = int(id)

		return b.Put(itob(t.ID), []byte(t.Description))
	})
}

// Remove method will delete the task from our db taking into account the id
// of the task
func (t *Task) Remove() error {
	return db.GetClient().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete(itob(t.ID))
	})
}

// List method will retrieve all the current tasks saved into our db
func List() (tasks []Task, err error) {
	err = db.GetClient().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				ID:          int(binary.BigEndian.Uint16(k)),
				Description: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
