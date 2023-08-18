package routers

import (
	"github.com/Zotish/Book-Store-Management-Porject-using-Golang-by-SQL/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterToBookStore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBOOK).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
