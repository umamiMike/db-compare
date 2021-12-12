package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
)

type dbStore struct {
	db *gorm.DB
}

func NewDB(dbName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	return db

}
func main() {

	db := NewDB("test.db")
	// db := dbStore.db
	defer db.Close()
	// Migrate the schema
	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}
	db.AutoMigrate(&Product{})
	fmt.Println("Product Table migrated")

	// products := db.Count(&Product{})
	// fmt.Printf("there are %v services" , products)

	if !db.HasTable(&Service{}) {
		db.CreateTable(&Service{})
	}

	fmt.Println("Service Table migrated")
	db.AutoMigrate(&Service{})

	var product Product
	var products []Product
	db.First(&product)
	//all products
	db.Find(&products)
	for i, e := range products {

}
	


	// services := db.Count(&Service{})
	// spew.Dump(services)
	// fmt.Printf("there are %v services" , services)

}
