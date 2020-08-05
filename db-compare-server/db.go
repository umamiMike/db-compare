/**
TODO: add option to start server with config file
EX: dbcompare --config /path/to/configfilename.ext */

package main

import (
	"database/sql"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func queryDb(data chan []map[string]interface{}, errorchan chan error, sqlString string, dbConn string) {

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		errorchan <- err
	}
	rows, err := db.Query(sqlString)
	if err != nil {
		errorchan <- err
		return
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		errorchan <- err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	db.Close()
	if err != nil {
		errorchan <- err
		return
	}
	data <- tableData // writing the data  to the data channel
}
