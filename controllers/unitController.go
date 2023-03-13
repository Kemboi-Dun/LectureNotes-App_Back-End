package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A unit
func UnitCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var unit_body struct {
		Name       string
		Code       string
		Course     string
		CourseType string
	}
	c.Bind(&unit_body)
	// CREATE A unit
	unit := models.Unit{
		Name:       unit_body.Name,
		Code:       unit_body.Code,
		Course:     unit_body.Course,
		CourseType: unit_body.CourseType,
	}
	result := initializers.DB.Create(&unit)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A unit
	c.JSON(http.StatusOK, gin.H{
		"unit": unit,
	})
}

// *!GET unitS
func GetUnits(c *gin.Context) {
	// GET unitS
	var units []models.Unit
	initializers.DB.Find(&units)
	// RESPOND WITH unit
	c.JSON(http.StatusOK, gin.H{
		"units": units,
	})
}

// *! EDIT unit
func GetUnitId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET unit
	var unit models.Unit
	initializers.DB.First(&unit, id)

	// RETURN UPDATED unit
	c.JSON(http.StatusOK, gin.H{
		"unit": unit,
	})
}

// *! UPDATE unit
func UpdateUnit(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var unit_body struct {
		Name       string
		Code       string
		Course     string
		CourseType string
	}
	c.Bind(&unit_body)

	// GET unit
	var unit models.Unit
	initializers.DB.First(&unit, id)
	// UPDATE unit
	initializers.DB.Model(&unit).Updates(models.Unit{
		Name:       unit_body.Name,
		Code:       unit_body.Code,
		Course:     unit_body.Course,
		CourseType: unit_body.CourseType,
	})
	// RETURN UPDATED unit
	c.JSON(http.StatusOK, gin.H{
		"unit": unit,
	})
}

// *! DELETE unit
func DeleteUnit(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET unit
	var unit models.Unit
	initializers.DB.Delete(&unit, id)

	// RETURN UPDATED unit
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
