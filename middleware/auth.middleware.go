package middleware

import (
	"github.com/gin-gonic/gin"
)

func Header() gin.HandlerFunc {
	return	func(c *gin.Context) {
		header := c.GetHeader("authorization")

		if header != "MyShop" {
			c.JSON(200,gin.H{
				"message": "No Connect",
			})
			c.Abort()
		}else {
			c.Next()
		}
		
	}
}