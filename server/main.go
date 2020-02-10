package main

import (
	//"encoding/csv"
	"encoding/json"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"github.com/recursionpharma/go-csv-map"
	"log"
	"net/http"
	"os"
)

const serverport = ":9099"

type conf struct {
	Dbs         []conndata `json:"dbs"`
	Server_port string     `json:"server_port"`
}

type conndata struct {
	Name string `json:"name"`
	Conn string `json:"conn"`
}

func (c *conf) get(filename string) *conf {
	fc, err := os.Open(filename)
	defer fc.Close()
	decoder := json.NewDecoder(fc)
	decoder.Decode(&c)
	if err != nil {
		log.Println(err)
	}
	return c

}

func csvThatStuff() ([]map[string]string, error) {
	recordFile, err := os.Open("/home/mike/google-drive/Documents/Financial/data/ynab_register.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil, err
	}

	// Setup the reader
	reader := csvmap.NewReader(recordFile)
	reader.Reader.LazyQuotes = true
	reader.Columns, err = reader.ReadHeader()

	if err != nil {
		fmt.Println(" error with ReadHeader", err)
		os.Exit(1)
	}
	// Read the records
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil, err
	}

	err = recordFile.Close()
	if err != nil {
		fmt.Println("An error encountered ::", err)
		return nil, err
	}

	//header := rawdat[0]
	//var dat []map[string]string

	//for rowindex, rowdata := range rows {
	//// the row becomes a map
	////{ "field": "value",

	////}
	//var xformed_row = map[string]string{}

	//for h, j := range rowdata {

	//fmt.Println(h, j)
	//}
	//append(dat, xformed_row)

	//}
	return allRecords, nil
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(serverport, nil)
}
