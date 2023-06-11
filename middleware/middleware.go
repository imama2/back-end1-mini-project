package middleware

import (
	"back-end1-mini-project/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
