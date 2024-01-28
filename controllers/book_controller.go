package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gowdaganesh005/MySQLserver/models"
	"github.com/gowdaganesh005/MySQLserver/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newbooks := models.GetAllBooks()
	res, err := json.Marshal(newbooks)
	if err != nil {
		fmt.Println("could not parse the data :", err)

	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookbyId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing:", err)
		bookdetails, _ := models.GetBookbyId(Id)
		res, _ := json.Marshal(bookdetails)
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	}
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseJson(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("could not parse the data :", err)

	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatebook = &models.Book{}
	utils.ParseJson(r, updatebook)

	id := chi.URLParam(r, "id")
	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("could not parse the data :", err)

	}
	bookdetails, db := models.GetBookbyId(Id)
	if bookdetails.Name != "" {
		bookdetails.Name = updatebook.Name
	}
	if bookdetails.Author != "" {
		bookdetails.Author = updatebook.Author
	}
	if bookdetails.Publication != "" {
		bookdetails.Publication = updatebook.Publication
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
