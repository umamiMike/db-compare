package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/icrowley/fake"
)

func main() {
	//sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	db, err := sql.Open("mysql", "")
	checkErr(err)
	rows, err := db.Query("SELECT id FROM patients")
	checkErr(err)
	// Get column names
	for rows.Next() {
		fname := fake.FirstName()
		lname := fake.LastName()
		mname := fake.Word()
		stmt, err := db.Prepare("Update patients set firstname=?, lastname =?, middlename=? WHERE id=?")
		checkErr(err)
		//made a var to hold the data
		var id int
		//scan the row for the appropriate column and assign, check for error and print out
		err = rows.Scan(&id)
		checkErr(err)
		res, err := stmt.Exec(fname, lname, mname, id)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(res)
	}
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

