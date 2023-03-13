package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A year
func YearCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var year_body struct {
		Name string
	}
	c.Bind(&year_body)
	// CREATE A year
	year := models.Year{
		Name: year_body.Name,
	}
	result := initializers.DB.Create(&year)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A year
	c.JSON(http.StatusOK, gin.H{
		"year": year,
	})
}

// *!GET yearS
func GetYears(c *gin.Context) {
	// GET yearS
	var years []models.Year
	initializers.DB.Find(&years)
	// RESPOND WITH year
	c.JSON(http.StatusOK, gin.H{
		"years": years,
	})
}

// *! EDIT year
func GetYearId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET year
	var year models.Year
	initializers.DB.First(&year, id)

	// RETURN UPDATED year
	c.JSON(http.StatusOK, gin.H{
		"year": year,
	})
}

// *! UPDATE year
func UpdateYear(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var year_body struct {
		Name string
	}
	c.Bind(&year_body)

	// GET year
	var year models.Year
	initializers.DB.First(&year, id)
	// UPDATE year
	initializers.DB.Model(&year).Updates(models.Year{
		Name: year_body.Name,
	})
	// RETURN UPDATED year
	c.JSON(http.StatusOK, gin.H{
		"year": year,
	})
}

// *! DELETE year
func DeleteYear(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET year
	var year models.Year
	initializers.DB.Delete(&year, id)

	// RETURN UPDATED year
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfully...",
	})
}
