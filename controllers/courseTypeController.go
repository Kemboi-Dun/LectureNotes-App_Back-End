package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A courseType
func CourseTypeCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var courseType_body struct {
		Name string
	}
	c.Bind(&courseType_body)
	// CREATE A courseType
	courseType := models.CourseType{
		Name: courseType_body.Name,
	}
	result := initializers.DB.Create(&courseType)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A courseType
	c.JSON(http.StatusOK, gin.H{
		"courseType": courseType,
	})
}

// *!GET courseTypeS
func GetCourseTypes(c *gin.Context) {
	// GET courseTypeS
	var courseTypes []models.CourseType
	initializers.DB.Find(&courseTypes)
	// RESPOND WITH courseType
	c.JSON(http.StatusOK, gin.H{
		"courseTypes": courseTypes,
	})
}

// *! EDIT courseType
func GetCourseTypeId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET courseType
	var courseType models.CourseType
	initializers.DB.First(&courseType, id)

	// RETURN UPDATED courseType
	c.JSON(http.StatusOK, gin.H{
		"courseType": courseType,
	})
}

// *! UPDATE courseType
func UpdateCourseType(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var courseType_body struct {
		Name string
	}
	c.Bind(&courseType_body)

	// GET courseType
	var courseType models.CourseType
	initializers.DB.First(&courseType, id)
	// UPDATE courseType
	initializers.DB.Model(&courseType).Updates(models.CourseType{
		Name: courseType_body.Name,
	})
	// RETURN UPDATED courseType
	c.JSON(http.StatusOK, gin.H{
		"courseType": courseType,
	})
}

// *! DELETE courseType
func DeleteCourseType(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET courseType
	var courseType models.CourseType
	initializers.DB.Delete(&courseType, id)

	// RETURN UPDATED courseType
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
