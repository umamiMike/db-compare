package main

import (
	"encoding/json"
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

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/datasource", datasourceHandler)
	http.ListenAndServe(serverport, nil)
}
