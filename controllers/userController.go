package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kemboi-dun/initializers"
	"github.com/kemboi-dun/models"
)

// *!CREATE A USER
func UserCreate(c *gin.Context) {
	// GET DATA OFF A REQUEST
	var user_body struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
		Role      string
	}
	c.Bind(&user_body)
	// CREATE A USER
	user := models.User{
		FirstName: user_body.FirstName,
		LastName:  user_body.LastName,
		Email:     user_body.Email,
		Password:  user_body.Password,
		Role:      user_body.Role,
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// RETURN A USER
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// *!GET USERS
func GetUsers(c *gin.Context) {
	// GET USERS
	var users []models.User
	initializers.DB.Find(&users)
	// RESPOND WITH USER
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// *! EDIT USER
func GetUserId(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET USER
	var user models.User
	initializers.DB.First(&user, id)

	// RETURN UPDATED USER
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// *! UPDATE USER
func UpdateUser(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")
	// GET DATA FROM REQ BODY
	var user_body struct {
		FirstName string
		LastName  string
		Email     string
		Password  string
		Role      string
	}
	c.Bind(&user_body)

	// GET USER
	var user models.User
	initializers.DB.First(&user, id)
	// UPDATE USER
	initializers.DB.Model(&user).Updates(models.User{
		FirstName: user_body.FirstName,
		LastName:  user_body.LastName,
		Email:     user_body.Email,
		Password:  user_body.Password,
		Role:      user_body.Role,
	})
	// RETURN UPDATED USER
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// *! DELETE USER
func DeleteUser(c *gin.Context) {
	// GET ID FROM URL
	id := c.Param("id")

	// GET USER
	var user models.User
	initializers.DB.Delete(&user, id)

	// RETURN UPDATED USER
	c.JSON(http.StatusOK, gin.H{
		"Message": "Record deleted succesfily...",
	})
}
