package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/umamimike/db-compare/db-compare-server/pkg/adding"
)

type msi = map[string]interface{}

//Convenience for typing
func Handler(a adding.Service) *chi.Mux {
	//setup routing and handlers
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request,norigin string) bool { return true }, AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", indexHandler)
	r.Get("/datasources", dsGetHandler(a))
	r.Post("/datasources", datasourcesPostHandler(a))

	return r

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	response := msi{
		"data": "stubby",
	}
	json.NewEncoder(w).Encode(response)

}

// ------------- Wrapper structs -----------------

type Data struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes msi    `json:"attributes"`
}

type DatasourceList struct {
	Type        string              `json:"type"`
	Datasources []adding.Datasource `json:"datasources"`
}

// ------------- datasource --------------------

func dsGetHandler(s adding.Service) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		dees := s.GetAll()
		resp := &DatasourceList{Type: "datasource", Datasources: dees}

		json.NewEncoder(w).Encode(resp)
	}
}

func datasourcesPostHandler(s adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var newDatasource adding.Datasource
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newDatasource)
		if err != nil {
			log.Println(err)
		}

		ID, _ := newUUID()
		resp := &Data{
			Id:   ID,
			Type: "datasource",
			Attributes: msi{
				"username": newDatasource.Username,
				"hostname": newDatasource.Hostname,
				"dbname":   newDatasource.DbName,
			},
		}
		s.AddDatasource(newDatasource)
		json.NewEncoder(w).Encode(resp)
	}
}
