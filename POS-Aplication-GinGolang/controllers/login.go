package controllers

import (
	"net/http"
	"os"
	"time"

	"vss/initializers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	initializers.LoadEnvVariables()

	// bind the json
	body := User{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	// connect to database
	var db, err = initializers.Connect()
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{
			"message": "cant connect to DB",
		})

		return
	}
	defer db.Close()

	// check Username
	var password *string

	err = db.
		QueryRow("SELECT password FROM users WHERE username = ?", body.Username).
		Scan(&password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username tidak terdaftar",
		})

		return
	}

	// check Password
	err = bcrypt.CompareHashAndPassword([]byte(*password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password salah",
		})

		return
	}

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute).Unix(),
		Issuer:    body.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cant generate jwt's token",
		})
	}

	// save cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": body.Username,
	})
}
