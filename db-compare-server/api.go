package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/umamimike/db-compare/db-compare-server/db"
)

var storageDir = "app-db-data/"

//Convenience for typing

func indexHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"data": "stub",
	}
	json.NewEncoder(w).Encode(response)

}

type msi = map[string]interface{}

// ------------- Wrapper structs -----------------

type Data struct {
	Id         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes map[string]interface{} `json:"attributes"`
}

// ------------- datasource --------------------

type DatasourceCredentials struct {
	Username string `json:"username"`
	Hostname string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

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

	// storage.Set(ID, response)
	// setJsonApiHeaders(w)
	rs, _ := json.Marshal(resp)
	fmt.Println(string(rs))
	w.Write(rs)
}

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
