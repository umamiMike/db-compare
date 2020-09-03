package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/umamimike/db-compare/db-compare-server/db"
)

var storageDir = "app-db-data/"

type msi = map[string]interface{}

//Convenience for typing

func indexHandler(w http.ResponseWriter, r *http.Request) {
	response := msi{
		"data": "stub",
	}
	json.NewEncoder(w).Encode(response)

}

// ------------- Wrapper structs -----------------

type Data struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes msi    `json:"attributes"`
}

// ------------- datasource --------------------

func datasourcesPostHandler(w http.ResponseWriter, r *http.Request) {

	storage := db.NewBadger(storageDir)
	defer storage.Close()

	var dsr DatasourceCredentials
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&dsr)

	ID, _ := newUUID()
	// passing the input data back to the output
	resp := &Data{
		Id:   ID,
		Type: "datasource",
		Attributes: msi{
			"username": dsr.Username,
			"hostname": dsr.Hostname,
			"dbname":   dsr.DbName,
		},
	}

	rs, _ := json.Marshal(resp)
	fmt.Println(string(rs))
	w.Write(rs)
}

func dsGetHandler() {}

func queriesPostHandler(w http.ResponseWriter, r *http.Request) {

	var query string
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&query)

	ID, _ := newUUID()
	// passing the input data back to the output
	response := &Data{
		Id:   ID,
		Type: "query",
		Attributes: msi{
			"query_string": query,
			"datasource":   "somedatasource",
		},
	}
	setJsonApiHeaders(w)
	json.NewEncoder(w).Encode(response)
}
