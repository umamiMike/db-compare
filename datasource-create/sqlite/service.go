package main

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model
	Name string
	Code string
	Cost uint

}
