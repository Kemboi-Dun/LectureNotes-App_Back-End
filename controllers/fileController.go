package controllers

import (
	"fmt"
	"log"
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

	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)

	// CREATE A FILE
	document := models.Documet{
		Name:       file.Filename,
		Path:       "assets/uploads/" + file.Filename,
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
	result := initializers.DB.Create(&document)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	// RETURN A FILE
	c.JSON(http.StatusOK, gin.H{
		"document": document,
	})
}

// *!GET FILES
func GetFiles(c *gin.Context) {
	// GET FILES
	var documents []models.Documet
	initializers.DB.Find(&documents)

	// RESPOND WITH FILE
	c.JSON(http.StatusOK, gin.H{
		"documents": documents,
	})
}

// *! GET DOCUMENT BY ID
func GetFileId(c *gin.Context) {

	// GET ID FROM URL
	id := c.Param("id")
	// GET FILE
	var document models.Documet
	initializers.DB.First(&document, id)

	// RETURN UPDATED FILE
	c.JSON(http.StatusOK, gin.H{
		"document": document,
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

	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// GET FILE
	var document models.Documet
	initializers.DB.First(&document, id)

	url := "localhost:3000/"

	// UPDATE FILE
	initializers.DB.Model(&document).Updates(models.Documet{
		Name:       file.Filename,
		Path:       url + "assets/uploads/" + file.Filename,
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
		"document": document,
	})
}

// *! DELETE FILE
func DeleteFile(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET file
	var document models.Documet
	initializers.DB.Delete(&document, id)

	// RETURN UPDATED FILE
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}

// func DownloadFile(c *gin.Context) {
// 	// Get the file ID or file name from the request parameters
// 	fileID := c.Param("fileId")

// 	// Open the file from the local file system
// 	filePath := "/assets/uploads/" + fileID
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer file.Close()

// 	// Set the response headers to indicate that a file is being downloaded
// 	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+fileID)
// 	c.Writer.Header().Set("Content-Type", "application/octet-stream")
// 	// c.Writer.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

// 	// Stream the file to the response
// 	io.Copy(c.Writer, file)
// }
