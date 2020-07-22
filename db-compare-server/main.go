package main

import (
	"fmt"
	"net/http"
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

func main() {
	// http.HandleFunc("/", indexHandler)
	http.HandleFunc("/datasources", datasourcesHandler)
	http.ListenAndServe(serverport, nil)
	fmt.Println("running on http://localhost:9099")
}
