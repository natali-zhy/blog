package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *sql.DB

func Connect() {

	conn, err := sql.Open("mysql", "root:@/article")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB OK")
	}
	DB = conn

}