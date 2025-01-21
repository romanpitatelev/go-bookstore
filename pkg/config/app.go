package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
	create a variable called db

(the whole point of app.go file is to return a variable db,
which will help the other files to intereact with db)
*/
var (
	db *gorm.DB
)

// Connect function helps to open connection with database (mysql database)
func Connect() {
	d, err := gorm.Open("mysql", "root:72328577Aa!@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
