package handler

import (
	"net/http"
	"strconv"

	"github.com/web-tuto-with-gin/model"
	"github.com/gin-gonic/gin"
)

// 홈 페이지 요청을 담당할 핸들러 함수
func ShowIndexPage(c *gin.Context) {
	articles := model.GetAllArticles()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func ShowArticleCreationPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}

// 글 조회 요청을 처리하는 핸들러 함수
func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := model.GetArticleByID(articleID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   article.Title,
				"payload": article}, "article.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// 글 생성 요청을 처리하는 핸들러 함수
func CreateArticle(c *gin.Context) {
	// Obtain the POSTed title and content values
	title := c.PostForm("title")
	content := c.PostForm("content")

	if _, err := model.CreateNewArticle(title, content); err == nil {
		// If the article is created successfully, redirect to home page
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		// if there was an error while creating the article, abort with an error
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

// 글 삭제 요청을 처리하는 핸들러 함수
func DeleteArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if err := model.DeleteArticleByID(articleID); err == nil {
			// If the article is deleted successfully, redirect to home page
			c.Redirect(http.StatusMovedPermanently, "/")
		} else {
			// if there was an error while deleting the article, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
