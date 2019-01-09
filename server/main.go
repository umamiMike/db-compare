package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	// "github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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

type myRespWriter struct {
	http.ResponseWriter
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
func (w *myRespWriter) writeHeader() {

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no config supplied")
		return
	}
	fmt.Println("service started")
	var cn conf
	cn.get(os.Args[1])
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+cn.Server_port, nil)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

type requestJson struct {
	QueryString string `json:"query"`
	DbKey       string `json:"db"`
	Token       string `json:"token"`
}

//func (w *LogResponseWritter) WriteHeader(statusCode int) {

//w.status = statusCode
//w.ResponseWriter.WriteHeader(statusCode)
//}
func indexHandler(w http.ResponseWriter, r *http.Request) {

	setHeaders(w)
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
		if v.Name == pd.DbKey {
			theconn = v.Conn
		}
	}
	fmt.Println(theconn)
	//os.Exit(1)
	datachannel := make(chan []map[string]interface{})
	errorchan := make(chan error)
	go getFromDb(datachannel, errorchan, pd.QueryString, theconn)

	select {
	case error := <-errorchan:
		errorStruct := struct {
			Token string `json:"token"`
			Dat   string `json: "err"`
		}{
			uuid,
			error.Error(),
		}
		jsonData, err := json.MarshalIndent(errorStruct, "", "  ")

		if err != nil {
			fmt.Println(err)

		}
		w.Write(jsonData)
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
		w.Write(jsonData)
	}

}
