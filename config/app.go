package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Connect() {
	conn, err := gorm.Open("mysql", "Ganesh:Ganesh@0011@/BookStore?charset=utf8&parseTime=True&loc=UTC")
	if err != nil {
		log.Fatal("Error opening the database :", err)
	}
	db = conn

}
func GetDB() *gorm.DB {
	return db
}
