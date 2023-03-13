package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A semester
func SemesterCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var semester_body struct {
		Course       string
		CourseType   string
		SemesterName string
	}
	c.Bind(&semester_body)
	// CREATE A semester
	semester := models.Semester{
		Course:       semester_body.Course,
		CourseType:   semester_body.CourseType,
		SemesterName: semester_body.SemesterName,
	}
	result := initializers.DB.Create(&semester)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A semester
	c.JSON(http.StatusOK, gin.H{
		"semester": semester,
	})
}

// *!GET semesterS
func GetSemesters(c *gin.Context) {
	// GET semesterS
	var semesters []models.Semester
	initializers.DB.Find(&semesters)
	// RESPOND WITH semester
	c.JSON(http.StatusOK, gin.H{
		"semesters": semesters,
	})
}

// *! EDIT semester
func GetSemesterId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET semester
	var semester models.Semester
	initializers.DB.First(&semester, id)

	// RETURN UPDATED semester
	c.JSON(http.StatusOK, gin.H{
		"semester": semester,
	})
}

// *! UPDATE semester
func UpdateSemester(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var semester_body struct {
		Course       string
		CourseType   string
		SemesterName string
	}
	c.Bind(&semester_body)

	// GET semester
	var semester models.Semester
	initializers.DB.First(&semester, id)
	// UPDATE semester
	initializers.DB.Model(&semester).Updates(models.Semester{
		Course:       semester_body.Course,
		CourseType:   semester_body.CourseType,
		SemesterName: semester_body.SemesterName,
	})
	// RETURN UPDATED semester
	c.JSON(http.StatusOK, gin.H{
		"semester": semester,
	})
}

// *! DELETE semester
func DeleteSemester(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET semester
	var semester models.Semester
	initializers.DB.Delete(&semester, id)

	// RETURN UPDATED semester
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
