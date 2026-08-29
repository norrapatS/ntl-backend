package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("access_token")

		if err != nil {
			authHeader := c.GetHeader("Authorization")

			if strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(jwtSecret), nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid or expired token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token claims",
			})
			c.Abort()
			return
		}

		userIDString, ok := claims["user_id"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid user id",
			})
			c.Abort()
			return
		}

		userID, err := uuid.Parse(userIDString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid user id",
			})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}