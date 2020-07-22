package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type requestJson struct {
	QueryString string `json:"query"`
	DbKey       string `json:"db"`
	Token       string `json:"token"`
}

func setJsonHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func setJsonApiHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/vnd.api+json")
}

//TODO: convert to chi, this should be named
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

// ------------- datasource --------------------

type DatasourceCredentials struct {
	Username string `json:"username"`
	Hostname string `json:"hostname"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func datasourcesHandler(w http.ResponseWriter, r *http.Request) {

	var dsr DatasourceCredentials
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&dsr)

	// passing the input data back to the output
	response := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "datasource",
			"id":   "1",
			"properties": map[string]interface{}{
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
