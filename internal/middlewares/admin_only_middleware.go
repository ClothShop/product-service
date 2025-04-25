package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, _ := c.Get("role")
		if userRole != "Admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Request forbidden"})
			return
		}

		c.Next()
	}
}
