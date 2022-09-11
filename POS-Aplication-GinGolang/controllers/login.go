package controllers

import (
	"fmt"
	"net/http"
	"time"

	"vss/conn"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *gin.Context) {
	mySigningKey := []byte("AllYourBase")

	// bind the json
	body := User{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// connect to database
	var db, err = conn.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// check username
	var id string
	err = db.
		QueryRow("SELECT id FROM users WHERE username = ?", body.Username).
		Scan(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username tidak terdaftar",
		})
		return
	}

	// check username and password
	err = db.
		QueryRow("SELECT id FROM users WHERE username = ? and password = ?", body.Username, body.Password).
		Scan(id)

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
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "cant generate token",
		})
	}

	// save cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"username": body.Username,
		"password": body.Password,
	})
}
