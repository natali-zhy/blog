package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_"github.com/gorilla/mux"
	"github.com/natalizhy/blog/pkg/db"
	"github.com/natalizhy/blog/pkg/models"
	"github.com/natalizhy/blog/utils"
	"net/http"
	"strconv"
)

var NewArticle models.Article

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	CreateArticle := &models.Article{}
	utils.ParseBody(r, CreateArticle)
	article, err := db.CreateArticle(*CreateArticle)
	fmt.Println(article, err)
	if err != nil {
		fmt.Println(err)
	}

	res, _ := json.Marshal(article)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	newArticle, err := db.GetAllArticle()
	fmt.Println(newArticle, err)
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
	schoolId := vars["articleId"]
	fmt.Println("schoolId", schoolId)
	ID, err := strconv.ParseInt(schoolId, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	schoolDetails, _ := db.GetArticleById(ID)
	res, _ := json.Marshal(schoolDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var updateArticle = &models.Article{}
	utils.ParseBody(r, updateArticle)
	vars := mux.Vars(r)
	schoolId := vars["articleId"]
	fmt.Println("schoolId", schoolId)
	ID, err := strconv.ParseInt(schoolId, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	fmt.Println("ID", ID)
	articleDetails, err := db.GetArticleById(ID)
	if err != nil {
		fmt.Println(err)
	}
	if updateArticle.Author != "" {
		articleDetails.Author = updateArticle.Author
	}
	if updateArticle.Text != "" {
		articleDetails.Text = updateArticle.Text
	}
	if updateArticle.Data != "" {
		articleDetails.Data = updateArticle.Data
	}
	//db.Save(&articleDetails)
	res, _ := json.Marshal(articleDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	schoolId := vars["schoolId"]
	ID, err := strconv.ParseInt(schoolId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	article := db.DeleteArticle(ID)
	res, _ := json.Marshal(article)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
