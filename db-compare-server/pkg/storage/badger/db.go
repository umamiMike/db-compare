package db

import (
	"log"

	"github.com/dgraph-io/badger"
)

type Badger struct {
	db *badger.DB
}

type Operation struct {
	Key   string
	Value []byte
	Op    byte
}

func NewBadger(storageDir string) *Badger {
	storage := &Badger{}
	opts := badger.DefaultOptions(storageDir)
	opts.SyncWrites = true
	opts.Dir = storageDir
	opts.ValueDir = storageDir
	opts.EventLogging = false
	var err error
	storage.db, err = badger.Open(opts)
	if err != nil {
		panic(err)
	}

	go storage.runStorageGC()

	return storage
}

// Set adds a key-value pair to the database
func (storage *Badger) Set(key string, value []byte) (err error) {
	return storage.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		return err
	})

}

func (db *Badger) Write(key string, data string) {
	err := db.Set(key, []byte(data))
	if err != nil {
		log.Fatal(err)
	}

}

// Del deletes a key
func (storage *Badger) Del(key string) (err error) {
	return storage.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
}

// Get returns value by key
func (storage *Badger) Get(key string) (value []byte, err error) {
	err = storage.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}
		return nil
	})
	return
}
