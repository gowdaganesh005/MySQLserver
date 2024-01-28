package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func connect() {
	conn, err := sql.Open("MySQL", "Ganesh:Ganesh@0011@/BookStore?charset=utf8&parseTime=True&loc=local")
	if err != nil {
		log.Fatal("Error opening the database :", err)
	}
	db = conn

}
func GetDB() *sql.DB {
	return db
}
