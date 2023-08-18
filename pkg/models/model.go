package models

import (
	"github.com/Zotish/Book-Store-Management-Porject-using-Golang-by-SQL/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `josn:"publication"`
}

func init() {
	config.Connect()
	db = config.GETDB()
	db.AutoMigrate(&Book{})
}
func (B *Book) CreateBook() *Book {
	db.NewRecord(B)
	db.Create(&B)
	return B
}
func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBook(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
func Delete(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
