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

//type JsonResponse struct:w

var (
	PORT = ":9099"
	DB   = map[string]string{
		"Test": "vagrant_test:vagrant@tcp(192.168.56.22:3306)/AllergyTest",
		"New":  "vagrant:vagrant@tcp(192.168.56.22:3306)/AllergyNew",
		"ae1":  "allergylocal:password@tcp(192.168.56.21:3336)/AllergyEMR",
		"v1":   "root:Ph7urIzJ5Gp2sfev@cloudsql(allergyedgev2:us-central1:temp-entaet)/AllergyEMR",
		"v2":   "root:Ph7urIzJ5Gp2sfev@cloudsql(allergyedgev2:us-central1:temp-entaet)/main",
	}
)

// Connection String details:
// * user      				 - the user created inside the DB. You can see more details on how to create it without password here:
//               			 https://cloud.google.com/sql/docs/sql-proxy#flags
// 										"user@cloudsql(project-id:zone:instance-name)/db-name?charset=utf8&parseTime=True&loc=UTC"
// CLOUD SQL FORMAT  	“user@cloudsql(project-id:zone:instance-name)/db”
//WHAT WILL GAVE ME 	-instances=allergyedgev2:us-central1:temp-entaet=tcp:3306 user: root pass: Ph7urIzJ5Gp2sfev
//
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
