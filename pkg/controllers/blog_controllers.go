package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/natalizhy/blog/pkg/db"
	"github.com/natalizhy/blog/pkg/models"
	"github.com/natalizhy/blog/utils"
	"net/http"
	"strconv"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	CreateArticle := &models.Article{}
	utils.ParseBody(r, CreateArticle)
	article, err := db.CreateArticle(*CreateArticle)
	if err != nil {
		fmt.Println(err)
	}

	res, _ := json.Marshal(article)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	newArticle, err := db.GetAllArticle()
	if err != nil {
		fmt.Println(err)
	}
	res, _ := json.Marshal(newArticle)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	ID, err := strconv.ParseInt(articleId, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	articleDetails, _ := db.GetArticleById(ID)
	res, _ := json.Marshal(articleDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var updateArticle = &models.Article{}
	utils.ParseBody(r, updateArticle)
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	ID, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	err = db.UpdateArticleById(ID, updateArticle)
	if err != nil {
		fmt.Println("Error UpdateArticleById")
	}

	articleDetails, err := db.GetArticleById(ID)
	if err != nil {
		fmt.Println("Error GetArticleById")
	}

	res, _ := json.Marshal(articleDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	articleId := vars["articleId"]
	ID, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	article := db.DeleteArticle(ID)
	res, _ := json.Marshal(article)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
