package adding

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
)

type Service interface {
	AddDatasource(...Datasource) error
	GetAll() []Datasource
}

//access to repository
type Repository interface {
	AddDatasource(Datasource) error
	GetAll() ([]string, error)
}
type service struct {
	r Repository
}

func NewService(r Repository) Service {
	log.Println("starting service")
	log.Println(r)
	return &service{r}
}

func (s *service) AddDatasource(ds ...Datasource) error {
	for _, d := range ds {
		fmt.Println("d in add datasource")
		spew.Dump(d)
		err := s.r.AddDatasource(d)
		if err != nil {
			log.Println("couldnt add datasource")
		}
	}
	return nil
}
func (s *service) AddList(dsl []Datasource) error {
	for _, d := range dsl {
		err := s.r.AddDatasource(d)
		if err != nil {
			log.Println("couldnt add datasource")
		}
	}
	return nil
}
func (s *service) GetAll() []Datasource {

	dsses, _ := s.r.GetAll()
	var outds []Datasource
	// converting to adding.Datasource for correct encoding
	for _, el := range dsses {
		var ds Datasource
		json.Unmarshal([]byte(el), &ds)
		outds = append(outds, ds)
	}

	return outds

}
