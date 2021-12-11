package dbconnection

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

//database connection defined seperate from main function for later use.
func Dbconnection() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=sifre.123 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	dbase := db.DB()
	err = dbase.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connection to database established.")
	}
	return db

}
