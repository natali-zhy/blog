package main

import (
	"github.com/gorilla/mux"
	"github.com/natalizhy/blog/pkg/db"
	"github.com/natalizhy/blog/routes"
	"log"
	"net/http"
)


func main() {
	db.Connect()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
