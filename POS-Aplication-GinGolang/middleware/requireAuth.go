package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"vss/initializers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {

	initializers.LoadEnvVariables()

	// get Cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cant get token from cookie",
		})

		return
	}

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	// logic authorization
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", claims["iss"])

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Next()
}
