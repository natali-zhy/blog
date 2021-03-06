package db

import (
	"fmt"
	"github.com/natalizhy/blog/pkg/models"
	"log"
)

func CreateArticle(article models.Article) (newUserId int64, err error) {
	newUserId = 0

	res, err := DB.Exec("INSERT INTO `Article` (`Author`, `Text`, `Data`) "+
		"VALUES (?, ?, ?)", article.Author, article.Text, article.Data)
	if err != nil {
		return
	}

	newUserId, err = res.LastInsertId()
	if err != nil {
		return
	}

	return
}

func GetAllArticle() (articles []models.Article, err error) {
	rows, err := DB.Query("SELECT `Id`, `Author`, `Text`, `Data` " +
		"FROM `Article`")
	if err != nil {
		return
	}

	for rows.Next() {
		article := models.Article{}
		err = rows.Scan(&article.Id, &article.Author, &article.Text, &article.Data)
		if err != nil {
			fmt.Println(err)
			return
		}
		articles = append(articles, article)
	}
	return
}
func GetArticleById(id int64) (article models.Article, err error) {
	err = DB.QueryRow("SELECT `Id`, `Author`, `Text`, `Data` "+
		"FROM `Article` WHERE id=?", id).Scan(&article.Id, &article.Author, &article.Text, &article.Data)

	if err != nil {
		return
	}

	return
}

func UpdateArticleById(ID int64, updateArticle *models.Article) (err error) {
	_, err = DB.Exec("UPDATE `Article` SET `Author`=?, `Text`=?, `Data`=? "+
		"WHERE `Id`=?", updateArticle.Author, updateArticle.Text, updateArticle.Data, ID)
	if err != nil {
		return
	}

	return
}

func DeleteArticle(ID int64) (err error) {
	_, err = DB.Exec("DELETE FROM `Article` "+
		"WHERE `id`=?", ID)

	if err != nil {
		log.Println(err)
	}
	return
}
