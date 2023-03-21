package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

func UploadDoc(c *gin.Context) {

	var doc_body struct {
		Path     string
		AuthorID int
	}
	c.Bind(&doc_body)

	// single file
	document, _ := c.FormFile("file")
	log.Println(document.Filename)

	// Upload the file to specific dst.
	c.SaveUploadedFile(document, "assets/uploads/"+document.Filename)

	file := models.Doc{
		Path:     "assets/uploads/" + document.Filename,
		AuthorID: doc_body.AuthorID,
	}
	result := initializers.DB.Create(&file)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", document.Filename))

	c.JSON(200, gin.H{
		"File": "assets/uploads/" + document.Filename,
	})
}

func GetDoc(c *gin.Context) {

	// GET FILES
	var files []models.Doc
	initializers.DB.Find(&files)
	// RESPOND WITH FILE
	c.JSON(http.StatusOK, gin.H{
		"Files": files,
	})

}

// *! EDIT FILE
func GetDocId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET FILE
	var document models.Doc
	initializers.DB.First(&document, id)

	// RETURN UPDATED FILE
	c.JSON(http.StatusOK, gin.H{
		"File": document,
	})
}

// *! UPDATE FILE
func UpdateDoc(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	var doc_body struct {
		Path     string
		AuthorID int
	}
	c.Bind(&doc_body)

	// single file
	document, _ := c.FormFile("file")
	log.Println(document.Filename)

	// GET FILE
	var file models.Doc
	initializers.DB.First(&file, id)
	// UPDATE USER
	initializers.DB.Model(&file).Updates(models.Doc{
		Path:     "assets/uploads/" + document.Filename,
		AuthorID: doc_body.AuthorID,
	})
	// RETURN UPDATED USER
	c.JSON(http.StatusOK, gin.H{
		"File": "assets/uploads/" + document.Filename,
	})
}

// *! DELETE USER
func DeleteDoc(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET FILE
	var file models.Doc
	initializers.DB.Delete(&file, id)

	// RETURN UPDATED USER
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfily...",
	})
}
