package controllers

import (
	"bookstore/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetBooksById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("BookID")
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	bookDetails, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	json.NewDecoder(r.Body).Decode(&createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	json.NewDecoder(r.Body).Decode(&UpdateBook)
	id := r.PathValue("BookID")
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	bookDetails, db := models.GetBookById(bookId)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("BookID")
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	book := models.DeleteBook(bookId)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
