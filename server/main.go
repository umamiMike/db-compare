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
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

var (
	PORT = ":9099"
	DB   = map[string]string{
		"whichdb": "user:password@tcp(ip.to.db:3306)/DBName",
	}
)

type PostData struct {
	QueryString string `json:"query"`
	DB          string `json: "db"`
}

func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
	var qs PostData
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&qs)
	dbData := getJSON(qs.QueryString, qs.DB)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(dbData)
}
func printLog(msg string, varToLog interface{}) {

	fmt.Print(msg + ": ")
	fmt.Println(varToLog)
}
func main() {
	if len(os.Args) < 2 {
		for key, _ := range DB {
			fmt.Println((key))
		}

	}
	fmt.Println("running server on " + PORT)
	http.HandleFunc("/", jsonResponseHandler)
	http.ListenAndServe(PORT, nil)
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getJSON(sqlString string, dbConn string) []byte {
	db, err := sql.Open("mysql", DB[dbConn])
	printLog("the dbconn in the getJSON is", DB[dbConn])
	checkErr(err)
	rows, err := db.Query(sqlString)
	if err != nil {
		var mt []byte
		return mt
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		var mt []byte
		return mt
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
	jsonData, err := json.MarshalIndent(tableData, "", "  ")
	if err != nil {
	}
	db.Close()
	return jsonData
}

func SaveLinkAsJson(l LinkInfo, dir string) {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir, 0755)
		} else {
			log.Println(err)
		}
	}

	path := fmt.Sprint(dir, url.QueryEscape(l.Link))
	os.Remove(path)

	b, err := json.Marshal(l)
	if err != nil {
		log.Println(err)
	}

	ioutil.WriteFile(path, b, 0644)
}
