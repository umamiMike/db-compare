package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/umamimike/db-compare/db-compare-server/pkg/adding"
	"github.com/umamimike/db-compare/db-compare-server/pkg/http/rest"
	"github.com/umamimike/db-compare/db-compare-server/pkg/storage/badger"
)

func main() {
	//storage temporarily hard coded as badger
	s, err := badger.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v storage ", s)
	var adder adding.Service

	server := &http.Server{
		Addr:         ":9099",
		Handler:      rest.Handler(adder),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	panic(server.ListenAndServe())

}

