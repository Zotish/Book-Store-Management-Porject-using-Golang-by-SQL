package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dataBase, err := gorm.Open("mysql", "SA:iLIKEYOU@1112@tcp(192.168.0.8)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db = dataBase
}
func GETDB() *gorm.DB {
	return db
}
