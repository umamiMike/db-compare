package json

import (
	"encoding/json"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/umamimike/db-compare/db-compare-server/pkg/adding"
	"github.com/umamimike/db-compare/db-compare-server/pkg/storage"
	"log"
	"path"
	"runtime"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionDatasource identifier for the JSON collection of beers
	CollectionDatasource = "datasource"
)

// Storage stores beer data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) AddDatasource(ds adding.Datasource) error {

	id, err := storage.GenID()
	if err != nil {
		log.Fatal(err)
	}

	newDS := Datasource{
		ID:       id,
		Username: ds.Username,
		Hostname: ds.Hostname,
		Password: ds.Password,
		DbName:   ds.DbName,
	}

	if err := s.db.Write(CollectionDatasource, id, newDS); err != nil {
		return err
	}
	return nil
}

func (s *Storage) GetAll() error {
	records, err := s.db.ReadAll(CollectionDatasource)
	log.Println(records)
	if err != nil {
		log.Println(err)

	}

	datasources := []Datasource{}
	for _, f := range records {
		dsFound := Datasource{}
		if err := json.Unmarshal([]byte(f), &dsFound); err != nil {
			log.Println(err)
		}
		datasources = append(datasources, dsFound)

	}
	log.Println(datasources)
	return nil
}
