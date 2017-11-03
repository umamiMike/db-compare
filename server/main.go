/**
TODO: convert the db map to an external file
change to parse that db file before any query attempt
so no need to restart.
TODO: ability to save slice of maps to external file
/config/save
will save the current config slice to an external file
/config/view
shows a web page of the current config
TODO: add option to start server with config file
EX: dbcompare --config /path/to/configfilename.ext
*/
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	//	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	//	"os"
)

var (
	PORT = ":9099"
)

type PostData struct {
	QueryString string `json:"query"`
	DbKey       string `json: "db"`
}

//call this function as a part of the process of finding what database to connect to
//will return a map of [string]strings
func getConn(cname string) string {
	data, err := ioutil.ReadFile("data.json")
	type Conf struct {
		Dbs  []map[string]string
		Port []map[string]string
	}
	conns := &Conf{}
	json.Unmarshal(data, &conns)
	if err != nil {
		fmt.Println("in getConn", err)
	}
	var thisconn string
	for _, num := range conns.Dbs {
		for k, v := range num {
			if k == cname {
				thisconn = v
			}
		}
	}
	fmt.Println("i am inside the getConn function", thisconn)
	return thisconn
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var qs PostData
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&qs)
	fmt.Println("in the index handler, and the db key is:", qs.DbKey)
	db := getConn(qs.DbKey)

	fmt.Println("the db is:", db)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	dbData, err := getJSON(qs.QueryString, db)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(dbData)
}

func main() {
	//	if len(os.Args) < 2 {
	//	}
	fmt.Println("running server on " + PORT)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(PORT, nil)
}

func getJSON(sqlString string, dbConn string) ([]byte, error) {
	db, err := sql.Open("mysql", dbConn)
	var mt []byte
	if err != nil {
		return mt, err
	}
	rows, err := db.Query(sqlString)
	if err != nil {
		return mt, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return mt, err
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
	jsonData, err := json.MarshalIndent(tableData, "", "  ")
	if err != nil {
		return mt, err
	}
	return jsonData, err
}
