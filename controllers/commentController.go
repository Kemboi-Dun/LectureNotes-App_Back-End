package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A comment
func CommentCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var comment_body struct {
		AuthorName string
		Body       string
		AuthorID   int
		FolderID   int
	}
	c.Bind(&comment_body)
	// CREATE A comment
	comment := models.Comment{
		AuthorName: comment_body.AuthorName,
		Body:       comment_body.Body,
		AuthorID:   comment_body.AuthorID,
		FolderID:   comment_body.FolderID,
	}
	result := initializers.DB.Create(&comment)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A comment
	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

// *!GET commentS
func GetComments(c *gin.Context) {
	// GET commentS
	var comments []models.Comment
	initializers.DB.Find(&comments)
	// RESPOND WITH comment
	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}

// *! EDIT comment
func GetCommentId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("folder_id")
	// GET comment
	var comment models.Comment
	initializers.DB.First(&comment, id)

	// RETURN UPDATED comment
	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

// *! UPDATE comment
func UpdateComment(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var comment_body struct {
		AuthorName string
		Body       string
		AuthorID   int
		FolderID   int
	}
	c.Bind(&comment_body)

	// GET comment
	var comment models.Comment
	initializers.DB.First(&comment, id)
	// UPDATE comment
	initializers.DB.Model(&comment).Updates(models.Comment{
		AuthorName: comment_body.AuthorName,
		Body:       comment_body.Body,
		AuthorID:   comment_body.AuthorID,
		FolderID:   comment_body.FolderID,
	})
	// RETURN UPDATED comment
	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

// *! DELETE comment
func DeleteComment(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET comment
	var comment models.Comment
	initializers.DB.Delete(&comment, id)

	// RETURN UPDATED comment
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
