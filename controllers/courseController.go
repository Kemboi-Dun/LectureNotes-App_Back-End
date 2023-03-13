package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A course
func CourseCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var course_body struct {
		Name string
	}
	c.Bind(&course_body)
	// CREATE A course
	course := models.Course{
		Name: course_body.Name,
	}
	result := initializers.DB.Create(&course)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A course
	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}

// *!GET courseS
func GetCourses(c *gin.Context) {
	// GET courseS
	var courses []models.Course
	initializers.DB.Find(&courses)
	// RESPOND WITH course
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}

// *! EDIT course
func GetCourseId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET course
	var course models.Course
	initializers.DB.First(&course, id)

	// RETURN UPDATED course
	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}

// *! UPDATE course
func UpdateCourse(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var course_body struct {
		Name string
	}
	c.Bind(&course_body)

	// GET course
	var course models.Course
	initializers.DB.First(&course, id)
	// UPDATE course
	initializers.DB.Model(&course).Updates(models.Course{
		Name: course_body.Name,
	})
	// RETURN UPDATED course
	c.JSON(http.StatusOK, gin.H{
		"course": course,
	})
}

// *! DELETE course
func DeleteCourse(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET course
	var course models.Course
	initializers.DB.Delete(&course, id)

	// RETURN UPDATED course
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
