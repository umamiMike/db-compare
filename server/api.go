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
	//querystring := pd.QueryString
	uuid := generateUUID(&pd)
	//var cn conf

	//datachannel := make(chan []map[string]interface{})
	//error_channel := make(chan error)

	//go getFromDb(datachannel, error_channel, querystring, theconn)

	//var resp interface{}
	//select {
	//case error := <-error_channel:
	//errorStruct := struct {
	//Token string `json:"token"`
	//Dat   string `json: "err"`
	//}{
	//uuid,
	//error.Error(),
	//}
	//jsonData, err := json.MarshalIndent(errorStruct, "", "  ")

	//if err != nil {
	//fmt.Println(err)

	//}

	//setHeaders(w)
	//w.Write(jsonData)

	//case message := <-datachannel:
	//messageStruct := struct {
	//Token string                   `json:"token"`
	//Dat   []map[string]interface{} `json: "dat"`
	//}{
	//uuid,
	//message,
	//}
	//jsonData, err := json.MarshalIndent(messageStruct, "", "  ")
	//if err != nil {
	//fmt.Println(err)
	//}

	rawdat, _ := csvThatStuff()
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
