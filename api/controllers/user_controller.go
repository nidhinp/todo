package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nidhinp/todo/api/models"
)

// Login endpoint to user login
func Login(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")

	login := models.User{}
	c.BindJSON(&login)
	validationErr := login.Validate("login")
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validationErr,
		})
		return
	}

	fmt.Println(login.Email)
	fmt.Println(login.Password)

	c.JSON(http.StatusOK, gin.H{
		"login": true,
	})
}
