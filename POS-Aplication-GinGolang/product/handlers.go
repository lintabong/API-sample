package product

import (
	"net/http"

	"vss/initializers"

	"github.com/gin-gonic/gin"
)

func GetSingleProduct(c *gin.Context) {

	initializers.LoadEnvVariables()

	id := c.Param("id")

	db, err := initializers.Connect()

	if err != nil {
		c.JSON(http.StatusFound, gin.H{
			"message": err.Error(),
		})

		return
	}
	defer db.Close()

	body := Product{}

	err = db.
		QueryRow("SELECT id, name, stock, sku, prize FROM products WHERE id = ?", id).
		Scan(&body.Id, &body.Name, &body.Stock, &body.SKU, &body.Prize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id not found",
		})

		return
	}

	c.JSON(http.StatusOK, body)
}

func UpdateProduct(c *gin.Context) {

	initializers.LoadEnvVariables()

	id := c.Param("id")

	// bind the json
	body := Product{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	db, err := initializers.Connect()

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})

		return
	}

	defer db.Close()

	_, err = db.Exec("update products set name=?, stock=?, SKU=?, prize=? where id=?",
		body.Name,
		body.Stock,
		body.SKU,
		body.Prize,
		id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, body)
}

func NewProduct(c *gin.Context) {

	initializers.LoadEnvVariables()

	body := Product{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	var db, err = initializers.Connect()

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})

		return
	}

	defer db.Close()

	_, err = db.Exec("insert into products (`name`, `stock`, `SKU`, `prize`) values (?, ?, ?, ?)",
		body.Name,
		body.Stock,
		body.SKU,
		body.Prize)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, body)

}

func DeleteProduct(c *gin.Context) {

	initializers.LoadEnvVariables()

	id := c.Param("id")

	var db, err = initializers.Connect()

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})

		return
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM products WHERE id=? ", id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success delete product",
	})
}
