package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A FOLDER
func FolderCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var folder_body struct {
		Name           string
		ParentFolderID int
	}
	c.Bind(&folder_body)
	// CREATE A FOLDER
	folder := models.Folder{
		Name:           folder_body.Name,
		ParentFolderID: folder_body.ParentFolderID,
	}
	result := initializers.DB.Create(&folder)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A FOLDER
	c.JSON(http.StatusOK, gin.H{
		"folder": folder,
	})
}

// *!GET FOLDERS
func GetFolders(c *gin.Context) {
	// GET FOLDERS
	var folders []models.Folder
	initializers.DB.Find(&folders)
	// RESPOND WITH FOLDER
	c.JSON(http.StatusOK, gin.H{
		"folders": folders,
	})
}

// *! EDIT FOLDER
func GetFolderId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET FOLDER
	var folder models.Folder
	initializers.DB.First(&folder, id)

	// RETURN UPDATED FOLDER
	c.JSON(http.StatusOK, gin.H{
		"folder": folder,
	})
}

// *! UPDATE FOLDER
func UpdateFolder(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var folder_body struct {
		Name           string
		ParentFolderID int
	}
	c.Bind(&folder_body)

	// GET FOLDER
	var folder models.Folder
	initializers.DB.First(&folder, id)
	// UPDATE FOLDER
	initializers.DB.Model(&folder).Updates(models.Folder{
		Name:           folder_body.Name,
		ParentFolderID: folder_body.ParentFolderID,
	})
	// RETURN UPDATED FOLDER
	c.JSON(http.StatusOK, gin.H{
		"folder": folder,
	})
}

// *! DELETE FOLDER
func DeleteFolder(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET FOLDER
	var folder models.Folder
	initializers.DB.Delete(&folder, id)

	// RETURN UPDATED FOLDER
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfily...",
	})
}
