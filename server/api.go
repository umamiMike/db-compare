package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
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

func indexHandler(w http.ResponseWriter, r *http.Request) {

	var pd requestJson
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&pd)
	uuid := generateUUID(&pd)

	// TODO: Convert to  recieving file upload????  url??? //
	rawdat, _ := parseCsvFile("/home/mike/google-drive/Documents/Financial/data/ynab_register.csv")
	//TODO: store in db???
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

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func generateUUID(r *requestJson) string {

	var uuid = ""
	if len(r.Token) <= 0 {
		uuid, _ = newUUID()
	}
	return uuid

}
