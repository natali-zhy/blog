package routes

import (
	"github.com/gorilla/mux"
	"github.com/natalizhy/blog/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/article/", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/article/", controllers.GetArticle).Methods("GET")
	router.HandleFunc("/article/{articleId}", controllers.GetArticleById).Methods("GET")
	router.HandleFunc("/article/{articleId}", controllers.UpdateArticle).Methods("PUT")
	router.HandleFunc("/article/{articleId}", controllers.DeleteArticle).Methods("DELETE")
}