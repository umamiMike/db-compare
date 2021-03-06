package json

import (
	"log"
	"path"
	"runtime"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/umamimike/db-compare/db-compare-server/pkg/adding"
	"github.com/umamimike/db-compare/db-compare-server/pkg/storage"
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

	if err := s.db.Write(CollectionDatasource, newDS.ID, ds); err != nil {
		return err
	}
	return nil
}
