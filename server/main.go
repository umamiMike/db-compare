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
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	// "github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
)

/** call this function as a part of the process of finding what database to connect to
* will return a map of [string]strings
searches config json in dbs function
*/
type conf struct {
	Dbs         []map[string]string `json:"dbs"`
	Server_port string              `json:"server_port"`
}
type requestJson struct {
	QueryString string `json:"query"`
	DbKey       string `json:"db"`
	Token       string `json:"token"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no config supplied")
		return
	}
	var cn conf
	cn.get(os.Args[1])
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+cn.Server_port, nil)
}

func (c *conf) get(f string) {
	fc, err := os.Open(f)
	defer fc.Close()
	decoder := json.NewDecoder(fc)
	decoder.Decode(&c)
	if err != nil {
		fmt.Println("in getConf", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am in the fn index handler")
	var pd requestJson
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&pd)

	var uuid = ""
	if len(pd.Token) <= 0 {
		uuid, _ = newUUID()
	}
	var cn conf
	cn.get(os.Args[1])
	var theconn = ""
	for _, v := range cn.Dbs {
		val, ok := v[pd.DbKey]
		if ok {
			theconn = val
		}
	}
	datachannel := make(chan []map[string]interface{})
	errorchan := make(chan error)

	go getJSON(datachannel, errorchan, pd.QueryString, theconn)

	select {

	case error := <-errorchan:
		fmt.Println(error, "will need to return an error")

	case message := <-datachannel:

		messageStruct := struct {
			Token string                   `json:"token"`
			Dat   []map[string]interface{} `json: "dat"`
		}{
			uuid,
			message,
		}

		jsonData, err := json.MarshalIndent(messageStruct, "", "  ")

		if err != nil {
			fmt.Println(err)

		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}

}
func getJSON(data chan []map[string]interface{}, errorchan chan error, sqlString string, dbConn string) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		errorchan <- err
	}
	rows, err := db.Query(sqlString)
	if err != nil {
		errorchan <- err
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
	}
	data <- tableData // writing the jsonData  to the data channel
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
