package main

import (
	"encoding/json"
	"os"
)

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
	}
	return c

}
