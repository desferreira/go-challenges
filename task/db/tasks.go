package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var tasksBucket = []byte("tasks")

var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

/*
Create a bucket to store the tasks
*/
func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucket)
		return err
	})
}

/*
Register a task on the bucket
*/
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := iotb(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

/*
Return all stored tasks
*/
func AllTasks() ([]Task, error) {
	var tasks []Task
	_ = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := Task{btoi(k), string(v)}
			tasks = append(tasks, task)
		}
		return nil
	})
	return tasks, nil
}

/*
Remove a task from the list
*/
func DeleteTask(num int) error {

	_ = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tasksBucket)
		err := b.Delete(iotb(num))
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

/*
Convert int to []byte
*/
func iotb(num int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(num))
	return b
}

/*
Convert []byte to int
*/
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
