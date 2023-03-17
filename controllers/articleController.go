package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A article
func ArticleCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var article_body struct {
		AuthorID   int
		Tag        string
		Title      string
		AuthorName string
		Body       string
	}
	c.Bind(&article_body)
	// CREATE A article
	article := models.Article{
		AuthorID:   article_body.AuthorID,
		Tag:        article_body.Tag,
		Title:      article_body.Title,
		AuthorName: article_body.AuthorName,
		Body:       article_body.Body,
	}
	result := initializers.DB.Create(&article)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A article
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

// *!GET articleS
func GetArticles(c *gin.Context) {
	// GET articleS
	var articles []models.Article
	initializers.DB.Find(&articles)
	// RESPOND WITH article
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

// *! EDIT article
func GetArticleId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET article
	var article models.Article
	initializers.DB.First(&article, id)

	// RETURN UPDATED article
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

// *! UPDATE article
func UpdateArticle(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var article_body struct {
		AuthorID   int
		Tag        string
		Title      string
		AuthorName string
		Body       string
	}
	c.Bind(&article_body)

	// GET article
	var article models.Article
	initializers.DB.First(&article, id)
	// UPDATE article
	initializers.DB.Model(&article).Updates(models.Article{
		AuthorID:   article_body.AuthorID,
		Tag:        article_body.Tag,
		Title:      article_body.Title,
		AuthorName: article_body.AuthorName,
		Body:       article_body.Body,
	})
	// RETURN UPDATED article
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
}

// *! DELETE article
func DeleteArticle(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET article
	var article models.Article
	initializers.DB.Delete(&article, id)

	// RETURN UPDATED article
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
