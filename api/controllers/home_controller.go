package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController return home response
func (s *Server) HomeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"home": "This is home controller",
	})
}
