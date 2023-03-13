package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A FILE
func FileCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var file_body struct {
		Name       string
		Path       string
		AuthorID   int
		FolderID   int
		Size       int64
		Year       int
		CourseType string
		CourseName string
		Semester   string
		UnitName   string
		UnitCode   string
		AuthorName string
		Type       string
	}
	c.Bind(&file_body)
	// CREATE A FILE
	file := models.File{
		Name:       file_body.Name,
		Path:       file_body.Path,
		AuthorID:   file_body.AuthorID,
		FolderID:   file_body.FolderID,
		Size:       file_body.Size,
		Year:       file_body.Year,
		CourseType: file_body.CourseType,
		CourseName: file_body.CourseName,
		Semester:   file_body.Semester,
		UnitName:   file_body.UnitName,
		UnitCode:   file_body.UnitCode,
		AuthorName: file_body.AuthorName,
		Type:       file_body.Type,
	}
	result := initializers.DB.Create(&file)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A FILE
	c.JSON(http.StatusOK, gin.H{
		"file": file,
	})
}

// *!GET FILES
func GetFiles(c *gin.Context) {
	// GET FILES
	var files []models.File
	initializers.DB.Find(&files)

	// RESPOND WITH FILE
	c.JSON(http.StatusOK, gin.H{
		"files": files,
	})
}

// *! EDIT file
func GetFileId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET FILE
	var file models.File
	initializers.DB.First(&file, id)

	// RETURN UPDATED FILE
	c.JSON(http.StatusOK, gin.H{
		"file": file,
	})
}

// *! UPDATE FILE
func UpdateFile(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var file_body struct {
		Name       string
		Path       string
		AuthorID   int
		FolderID   int
		Size       int64
		Year       int
		CourseType string
		CourseName string
		Semester   string
		UnitName   string
		UnitCode   string
		AuthorName string
		Type       string
	}
	c.Bind(&file_body)

	// GET FILE
	var file models.File
	initializers.DB.First(&file, id)
	// UPDATE FILE
	initializers.DB.Model(&file).Updates(models.File{
		Name:       file_body.Name,
		Path:       file_body.Path,
		AuthorID:   file_body.AuthorID,
		FolderID:   file_body.FolderID,
		Size:       file_body.Size,
		Year:       file_body.Year,
		CourseType: file_body.CourseType,
		CourseName: file_body.CourseName,
		Semester:   file_body.Semester,
		UnitName:   file_body.UnitName,
		UnitCode:   file_body.UnitCode,
		AuthorName: file_body.AuthorName,
		Type:       file_body.Type,
	})
	// RETURN UPDATED FILE
	c.JSON(http.StatusOK, gin.H{
		"file": file,
	})
}

// *! DELETE FILE
func DeleteFile(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET file
	var file models.File
	initializers.DB.Delete(&file, id)

	// RETURN UPDATED FILE
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
