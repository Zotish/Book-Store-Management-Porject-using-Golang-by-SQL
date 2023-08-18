package main

import (
	"log"
	"net/http"

	"github.com/Zotish/Book-Store-Management-Porject-using-Golang-by-SQL/pkg/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterToBookStore(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
