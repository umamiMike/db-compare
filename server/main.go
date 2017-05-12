package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//type JsonResponse struct:w

var (
	PORT = ":9099"
	DB   = map[string]string{
		"Test": "vagrant_test:vagrant@tcp(192.168.56.22:3306)/AllergyTest",
		"New":  "vagrant:vagrant@tcp(192.168.56.22:3306)/AllergyNew",
		"ae1":  "allergylocal:password@tcp(192.168.56.21:3336)/AllergyEMR",
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
	fmt.Println(qs.QueryString)
	fmt.Println(qs.DB)
	dbData := getJSON(qs.QueryString, qs.DB)
	w.Header().Set("Content-Type", "application/json")
	resp := json.NewEncoder(w).Encode(dbData)
	fmt.Println(resp)
	fmt.Fprint(w, resp)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: aecompare  'Name of Database you want to connect to, 'sql query wrapped in double quotes")
		fmt.Println("Databases Names available are")
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
	// db, err := sql.Open("mysql", "vagrant:vagrant@tcp(192.168.56.22:3306)/AllergyNew")
	db, err := sql.Open("mysql", DB[dbConn])
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
