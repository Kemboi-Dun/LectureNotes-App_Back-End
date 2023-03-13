package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A semesterName
func SemesterNameCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var semesterName_body struct {
		Name string
	}
	c.Bind(&semesterName_body)
	// CREATE A semesterName
	semesterName := models.SemesterName{
		Name: semesterName_body.Name,
	}
	result := initializers.DB.Create(&semesterName)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A semesterName
	c.JSON(http.StatusOK, gin.H{
		"semesterName": semesterName,
	})
}

// *!GET semesterNameS
func GetSemesterNames(c *gin.Context) {
	// GET semesterNameS
	var semesterNames []models.SemesterName
	initializers.DB.Find(&semesterNames)
	// RESPOND WITH semesterName
	c.JSON(http.StatusOK, gin.H{
		"semesterNames": semesterNames,
	})
}

// *! EDIT semesterName
func GetSemesterNameId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET semesterName
	var semesterName models.SemesterName
	initializers.DB.First(&semesterName, id)

	// RETURN UPDATED semesterName
	c.JSON(http.StatusOK, gin.H{
		"semesterName": semesterName,
	})
}

// *! UPDATE semesterName
func UpdateSemesterName(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var semesterName_body struct {
		Name string
	}
	c.Bind(&semesterName_body)

	// GET semesterName
	var semesterName models.SemesterName
	initializers.DB.First(&semesterName, id)
	// UPDATE semesterName
	initializers.DB.Model(&semesterName).Updates(models.SemesterName{
		Name: semesterName_body.Name,
	})
	// RETURN UPDATED semesterName
	c.JSON(http.StatusOK, gin.H{
		"semesterName": semesterName,
	})
}

// *! DELETE semesterName
func DeleteSemesterName(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET semesterName
	var semesterName models.SemesterName
	initializers.DB.Delete(&semesterName, id)

	// RETURN UPDATED semesterName
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
