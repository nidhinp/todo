package middlewares

import (
	"github.com/gin-gonic/gin"
)

// SetJSONMiddleware make the response a json
func SetJSONMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		next(c)
	}
}
