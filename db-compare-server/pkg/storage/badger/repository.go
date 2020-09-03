package badger

import (
	"log"

	"encoding/json"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/umamimike/db-compare/db-compare-server/pkg/adding"
	"github.com/umamimike/db-compare/db-compare-server/pkg/storage"
)

type Storage struct {
	db *badger.DB
}

type Operation struct {
	Key   string
	Value []byte
	Op    byte
}

func NewStorage(storageDir string) (*Storage, error) {
	storage := &Storage{}
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

	return storage, nil
}

func (s *Storage) AddDatasource(ds adding.Datasource) error {
	id, err := storage.GenID()
	if err != nil {
		log.Fatal("could not generate id", err)
	}

	newDS := Datasource{
		ID:       id,
		Username: ds.Username,
		Hostname: ds.Hostname,
		Password: ds.Password,
		DbName:   ds.DbName,
	}
	save, _ := json.Marshal(newDS)
	s.Write(id, save)
	return nil
}

//====================== boilerplate =========================================//
// Set adds a key-value pair to the database
func (storage *Storage) Set(key string, value []byte) (err error) {
	return storage.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), value)
		return err
	})

}

func (db *Storage) Write(key string, data []byte) {

	err := db.Set(key, data)
	if err != nil {
		log.Fatal(err)
	}

}

// Del deletes a key
func (storage *Storage) Del(key string) (err error) {
	return storage.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
}

// Get returns value by key
func (storage *Storage) Get(key string) (value []byte, err error) {
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

func (storage *Storage) runStorageGC() {
	timer := time.NewTicker(10 * time.Minute)
	for {
		select {
		case <-timer.C:
			storage.storageGC()
		}
	}
}
func (storage *Storage) storageGC() {
again:
	err := storage.db.RunValueLogGC(0.5)
	if err == nil {
		goto again
	}

}

func (storage *Storage) Close() {
	storage.db.Close()
}
