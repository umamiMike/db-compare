package main

import (
	"testing"
	"github.com/jinzhu/gorm"
)


func NewStore(dbName string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		panic("failed to connect database")
	}
	return db

}

func TestGet(t *testing.T) {
	// Create
	db.Create(&Product{Code: "L1212", Cost: 1000})

}
