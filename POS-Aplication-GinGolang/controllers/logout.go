package controllers

import (
	"net/http"
	"vss/initializers"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	initializers.LoadEnvVariables()

	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "already loged out",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "logout success",
	})
}
