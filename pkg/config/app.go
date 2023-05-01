// This file contains all the DB connection details

package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // this underscore infront of the package name is to import the side effects of this package without actaully importing it

	//call the inut function for this package
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error while gettting env file")
	}
	return os.Getenv(key)
}

func Connect() {
	d, err := gorm.Open("mysql", getEnv("ROOT_USERNAME")+":"+getEnv("ROOT_PASSWORD")+"@tcp(0.0.0.0:3306)/"+getEnv("DATABASE_NAME")+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
