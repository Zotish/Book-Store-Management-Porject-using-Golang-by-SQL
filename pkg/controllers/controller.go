package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zotish/Book-Store-Management-Porject-using-Golang-by-SQL/pkg/models"
	"github.com/Zotish/Book-Store-Management-Porject-using-Golang-by-SQL/pkg/uilts"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBOOK(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	BookId := params["BookId"]
	ID, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	BookDetails, _ := models.GetBook(ID)
	res, _ := json.Marshal(BookDetails)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	uilts.ParseBody(r, CreateBook)
	B := CreateBook.CreateBook()
	res, _ := json.Marshal(B)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	uilts.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookID := params["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBook(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Name = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(&bookDetails)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	B := models.Delete(ID)
	res, _ := json.Marshal(B)
	w.Header().Set("Content-Type", "publication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
