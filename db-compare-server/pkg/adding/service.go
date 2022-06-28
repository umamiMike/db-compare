package adding

import (
	"log"
)

type Service interface {
	AddDatasource(...Datasource) error
	GetAll() error
}

//access to repository
type Repository interface {
	AddDatasource(Datasource) error
	GetAll() error
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
func (s *service) GetAll() error {
s.r.GetAll()
// log.Println("view all datasources")

return nil
}
