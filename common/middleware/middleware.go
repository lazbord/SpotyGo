package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	err := VerifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Token verification failed": err})
		c.Abort()
	}
	c.Next()
}

func VerifyToken(c *gin.Context) error {
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		return errors.New("no token")
	}

	key := os.Getenv("JWT_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), err
	})

	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("not a valid token")
	}

	return nil
}
