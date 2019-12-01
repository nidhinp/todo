package controllers

import (
	"github.com/gin-gonic/gin"
)

// HomeController return home response
func HomeController(c *gin.Context) {
	c.JSON(200, gin.H{
		"home": "This is home controller",
	})
}
