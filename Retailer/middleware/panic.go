package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// PanicRecovery is a middleware to gracefully recover from panics.
func PanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic recovered: %v", r)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error. The service has recovered."})
				c.Abort()
			}
		}()
		c.Next()
	}
}
