package models

import (
	"github.com/gowdaganesh005/MySQLserver/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"date"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b

}
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books

}
func GetBookbyId(id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("ID=?", id).Find(getbook)
	return &getbook, db

}
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book

}
