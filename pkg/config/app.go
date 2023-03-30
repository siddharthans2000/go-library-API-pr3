// This file contains all the DB connection details

package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // this underscore infront of the package name is to import the side effects of this package without actaully importing it
	//call the inut function for this package
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "development:Develop@tcp(localhost:3306)/book_library?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
