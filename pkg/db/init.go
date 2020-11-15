package db

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB


func New(user string, password string, addr string, database string) {
	dataSourceName := user + ":" + password + "@tcp(" + addr + ")/" + database + "?parseTime=true"

	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Errorf("db: %w", err)
	}

	if err = db.Ping(); err != nil {
		fmt.Errorf("db: %w", err)
	}

	DB = db
}
