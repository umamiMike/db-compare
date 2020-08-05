package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Convenience for typing
type msi = map[string]interface{}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{
		"data": "stub",
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	setJsonHeaders(w)
	w.Write(jsonData)

}

// ------------- JsonApi structs -----------------
type JsonApi struct {
	Data Data `json: "data"`
}
type Data struct {
	Id         string `json: "id"`
	Type       string `json: "type"`
	Attributes msi    `json: "attributes"`
}

// ------------- datasource --------------------

type DatasourceCredentials struct {
	Username string `json:"username"`
	Hostname string `json:"host"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func datasourcesHandler(w http.ResponseWriter, r *http.Request) {

	var dsr DatasourceCredentials
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&dsr)

	// passing the input data back to the output
	response := JsonApi{
		Data: Data{
			Type: "datasource",
			Id:   "1",
			Attributes: msi{
				"username": dsr.Username,
				"host":     dsr.Hostname,
				"password": dsr.Password,
				"dbname":   dsr.DbName,
			},
		},
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	setJsonApiHeaders(w)
	w.Write(jsonData)
}

func queriesPostHandler(w http.ResponseWriter, r *http.Request) {

	var query string
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&query)

	// passing the input data back to the output
	response := JsonApi{
		Data: Data{
			Type: "query",
			Id:   "1",
			Attributes: msi{
				"query_string": query,
				"datasource":   "somedatasource",
			},
		},
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	setJsonApiHeaders(w)
	w.Write(jsonData)
}
