package main

import (
	"github.com/go-chi/chi"
	"github.com/natalizhy/blog/pkg/controllers"
	"github.com/natalizhy/blog/pkg/db"
	"log"
	"net/http"
)


func main() {
	r := chi.NewRouter()
	r.Get("/schools/", controllers.GetArticle)
	r.Post("/schools/", controllers.CreateArticle)
	r.Get("/schools/{articleId}", controllers.GetArticleById)
	r.Put("/schools/{articleId}", controllers.UpdateArticle)
	r.Delete("/schools/{schoolId}", controllers.DeleteArticle)
	db.Connect()

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
