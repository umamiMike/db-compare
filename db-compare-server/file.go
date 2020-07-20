package main

import (
	"encoding/json"
	"log"
	"os"
)

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
