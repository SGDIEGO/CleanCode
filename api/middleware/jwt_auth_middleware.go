package middleware

import (
	"net/http"

	"github.com/SGDIEGO/CleanCode/pkg/util"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secretKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {

		err := util.TokenValid(c, secretKey)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"errorAuth": "no auth",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
