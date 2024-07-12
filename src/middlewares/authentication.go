package middlewares

import (
	"github.com/BinDruid/spotify-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
