package middleware

import "github.com/gin-gonic/gin"

// This is one request middleware of name authenticate
func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"code":    401,
			"Message": "Token not present",
		})
		return
	}
	c.Next()
}

// another way of middleware
//func Authenticate() gin.HandlerFunc {
//	//write custom logic to be applied before my middleware is executed
//	return func(c *gin.Context) {
//		if !(c.Request.Header.Get("Token") == "auth") {
//			c.AbortWithStatusJSON(500, gin.H{
//				"code":    401,
//				"Message": "Token not present",
//			})
//			return
//		}
//		c.Next()
//	}
//}

func Addheader(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	c.Next()
}
