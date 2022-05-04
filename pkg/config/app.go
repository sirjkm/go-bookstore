package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbo * gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", //"name:password/tablename?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}