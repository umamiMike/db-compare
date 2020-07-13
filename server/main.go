package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	csvmap "github.com/recursionpharma/go-csv-map"
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

// called from api indexHandler

func parseCsvFile(csvFile string) ([]map[string]string, error) {

	recordFile, err := os.Open(csvFile)
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

	return allRecords, nil
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/datasource", datasourceHandler)
	http.ListenAndServe(serverport, nil)
}
