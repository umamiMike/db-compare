package adding

import (
	// "github.com/umamimike/db-compare/db-compare-server/pkg/adding"
	"log"
)

type Service interface {
	AddDatasource(...Datasource) error
	AddList([]Datasource) error
}

//access to repository
type Repository interface {
	AddDatasource(Datasource) error
}
type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddDatasource(ds ...Datasource) error {
	for _, d := range ds {
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
