package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A FILEACCESS
func FileAccessCreate(c *gin.Context) {

	// GET DATA OFF A REQUEST
	var fileAccess_body struct {
		FileID      int
		UserID      int
		AccessLevel string
	}
	c.Bind(&fileAccess_body)

	// CREATE A FILEACCESS
	fileAccess := models.FileAccess{
		FileID:      fileAccess_body.FileID,
		UserID:      fileAccess_body.UserID,
		AccessLevel: fileAccess_body.AccessLevel,
	}
	result := initializers.DB.Create(&fileAccess)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// RETURN A FILEACCESS
	c.JSON(http.StatusOK, gin.H{
		"fileAccess": fileAccess,
	})
}

// *!GET FILEACCESS
func GetFileAccess(c *gin.Context) {
	// GET FILEACCESSS
	var fileAccess []models.FileAccess
	initializers.DB.Find(&fileAccess)

	// RESPOND WITH fileAccess
	c.JSON(http.StatusOK, gin.H{
		"fileAccess": fileAccess,
	})
}

// *! EDIT FILEACCESS
func GetFileAccessId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET FILEACCESS
	var fileAccess models.FileAccess
	initializers.DB.First(&fileAccess, id)

	// RETURN UPDATED FILEACCESS
	c.JSON(http.StatusOK, gin.H{
		"fileAccess": fileAccess,
	})
}

// *! UPDATE FILEACCESS
func UpdateFileAccess(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var fileAccess_body struct {
		FileID      int
		UserID      int
		AccessLevel string
	}
	c.Bind(&fileAccess_body)

	// GET FILEACCESS
	var fileAccess models.FileAccess
	initializers.DB.First(&fileAccess, id)

	// UPDATE FILEACCESS
	initializers.DB.Model(&fileAccess).Updates(models.FileAccess{
		FileID:      fileAccess_body.FileID,
		UserID:      fileAccess_body.UserID,
		AccessLevel: fileAccess_body.AccessLevel,
	})

	// RETURN UPDATED FILEACCESS
	c.JSON(http.StatusOK, gin.H{
		"fileAccess": fileAccess,
	})
}

// *! DELETE FILEACCESS
func DeleteFileAccess(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET FILEACCESS
	var fileAccess models.FileAccess
	initializers.DB.Delete(&fileAccess, id)

	// RETURN UPDATED FILEACCESS
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfily...",
	})
}
