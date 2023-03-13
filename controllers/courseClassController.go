package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A courseClass
func CourseClassCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var courseClass_body struct {
		Year     string
		Semester string
	}
	c.Bind(&courseClass_body)
	// CREATE A courseClass
	courseClass := models.CourseClass{
		Year:     courseClass_body.Year,
		Semester: courseClass_body.Semester,
	}
	result := initializers.DB.Create(&courseClass)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A courseClass
	c.JSON(http.StatusOK, gin.H{
		"courseClass": courseClass,
	})
}

// *!GET courseClasses
func GetCourseClasses(c *gin.Context) {
	// GET courseClasses
	var courseClasses []models.CourseClass
	initializers.DB.Find(&courseClasses)
	// RESPOND WITH courseClass
	c.JSON(http.StatusOK, gin.H{
		"courseClasses": courseClasses,
	})
}

// *! EDIT courseClass
func GetCourseClassId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET courseClass
	var courseClass models.CourseClass
	initializers.DB.First(&courseClass, id)

	// RETURN UPDATED courseClass
	c.JSON(http.StatusOK, gin.H{
		"courseClass": courseClass,
	})
}

// *! UPDATE courseClass
func UpdateCourseClass(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var courseClass_body struct {
		Year     string
		Semester string
	}
	c.Bind(&courseClass_body)

	// GET courseClass
	var courseClass models.CourseClass
	initializers.DB.First(&courseClass, id)
	// UPDATE courseClass
	initializers.DB.Model(&courseClass).Updates(models.CourseClass{
		Year:     courseClass_body.Year,
		Semester: courseClass_body.Semester,
	})
	// RETURN UPDATED courseClass
	c.JSON(http.StatusOK, gin.H{
		"courseClass": courseClass,
	})
}

// *! DELETE courseClass
func DeleteCourseClass(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET courseClass
	var courseClass models.CourseClass
	initializers.DB.Delete(&courseClass, id)

	// RETURN UPDATED courseClass
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
