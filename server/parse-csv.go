package main

import (
	"fmt"
	"os"

	csvmap "github.com/recursionpharma/go-csv-map"
)

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
