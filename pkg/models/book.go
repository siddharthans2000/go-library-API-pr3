package models

import (
	"go-library-API-pr3/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})                         // User to create a schema with the book object
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)                                  // this condition takes all the values from the db and store it in the Books variable.
	return Books
}
func GetBookById(id int64) (*Book,*gorm.DB){
	var b Book
	db:=db.Where("ID=?",id).Find(&b)
	return &b,db
}
func DeleteById(id int64) Book{
	var b Book
	db.Where("ID=?",id).Delete(b)
	return b
}