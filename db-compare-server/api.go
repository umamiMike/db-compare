package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
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

//TODO: convert to chi, this should be named
func indexHandler(w http.ResponseWriter, r *http.Request) {

	var pd requestJson
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&pd)
	uuid := generateUUID(&pd)

	// TODO: Convert to  recieving file upload????  url??? //
	rawdat, _ := parseCsvFile("/home/mike/google-drive/Documents/Financial/data/ynab_register.csv")
	//TODO: store in db???
	response := map[string]interface{}{
		"index": "silence is golden",
	}	setJsonHeaders(w)
	w.Write(jsonData)
}

// ------------- datasource --------------------

type datasourceConnInfo struct {
	Username string `json:"username"`
	Hostname string `json:"hostname"`
	Password string `json:"password"`
	DbName   string `json:"db-name"`
}

func datasourceHandler(w http.ResponseWriter, r *http.Request) {

	var dsr datasourceConnInfo
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&dsr)

	spew.Dump(dsr)

	// passing the input data back to the output
	response := map[string]interface{}{
		"username": dsr.Username,
		"host":     dsr.Hostname,
		"password": dsr.Password,
		"dbname":   dsr.DbName,
	}

	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	setJsonHeaders(w)
	w.Write(jsonData)
}

func generateUUID(r *requestJson) string {
	var uuid = ""
	if len(r.Token) <= 0 {
		uuid, _ = newUUID()
	}
	return uuid
}
