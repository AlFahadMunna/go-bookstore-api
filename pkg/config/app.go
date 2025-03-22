package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/grom/dialects/mysql"
)

var db (*gorm.DB)

func Connect() {
	d, err := gorm.Open("mysql", "fahad:jahad@12/book-api?charser=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
