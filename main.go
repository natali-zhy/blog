package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natalizhy/blog/pkg/db"
	"github.com/natalizhy/blog/routes"
	"log"
	"net/http"
	"os"
)

type Application struct {
	AddrHttp         string
	Templates        string
	Db               Db
}

type Db struct {
	AddrMysql     string
	UserMysql     string
	PasswordMysql string
	Database      string
}

var config = &Application{}

func main() {

	file, err := os.Open("../blog/pkg/config/config.json")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB OK")
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		fmt.Println(err)
	}

	 db.New(config.Db.UserMysql, config.Db.PasswordMysql, config.Db.AddrMysql, config.Db.Database)

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(config.AddrHttp, r))

}
