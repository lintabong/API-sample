package controllers

import (
	"fmt"
	"net/http"
	"vss/initializers"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

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
		fmt.Println(err.Error())

		return
	}
	defer db.Close()

	pass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 2)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "cant hash the password",
		})

		return
	}

	//insert data to database
	_, err = db.Exec("INSERT INTO users (`username`, `password`, `contact`, `account_type`) VALUES (?, ?, ?, ?)",
		body.Username,
		pass,
		body.Contact,
		body.Account_type)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Register berhasil",
		})
	}

}
