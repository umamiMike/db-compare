package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type myRespWriter struct {
	http.ResponseWriter
}

type requestJson struct {
	QueryString string `json:"query"`
	DbKey       string `json:"db"`
	Token       string `json:"token"`
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func generateUUID(r *requestJson) string {

	var uuid = ""
	if len(r.Token) <= 0 {
		uuid, _ = newUUID()
	}
	return uuid

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	var pd requestJson
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&pd)
	uuid := generateUUID(&pd)

	rawdat, _ := parseCsvFile("/home/mike/google-drive/Documents/Financial/data/ynab_register.csv")
	messageStruct := struct {
		Token string      `json:"token"`
		Data  interface{} `json: "dat"`
	}{
		uuid,
		rawdat,
	}
	jsonData, err := json.MarshalIndent(messageStruct, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	setHeaders(w)
	w.Write(jsonData)
}
