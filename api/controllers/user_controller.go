package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Login struct
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate the login struct
func (login Login) Validate() error {
	return validation.ValidateStruct(&login,
		validation.Field(&login.Email, validation.Required, is.Email),
		validation.Field(&login.Password, validation.Required),
	)
}

// Login endpoint to user login
func (s *Server) Login(c *gin.Context) {
	var login Login
	c.BindJSON(&login)

	err := login.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"login": true,
	})
}
