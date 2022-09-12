package product

import (
	"fmt"
	"net/http"

	"vss/initializers"

	"github.com/gin-gonic/gin"
)

type Product struct {
	id    int
	name  string
	stock int
	SKU   string
	prize int
}

func GetSingleProduct(c *gin.Context) {

	id := c.Param("id")
	var db, err = initializers.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var prod = Product{}

	err = db.
		QueryRow("select id, name, stock, sku, prize from products where id = ?", id).
		Scan(&prod.id, &prod.name, &prod.stock, &prod.SKU, &prod.prize)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    prod.id,
		"name":  prod.name,
		"stock": prod.stock,
		"SKU":   prod.SKU,
		"prize": prod.prize,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	stock := c.Query("stock")
	SKU := c.Query("SKU")
	prize := c.Query("prize")

	var db, err = initializers.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("update products set name=?, stock=?, SKU=?, prize=? where id=?",
		name,
		stock,
		SKU,
		prize,
		id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"stock": stock,
		"SKU":   SKU,
		"prize": prize,
	})
}

func NewProduct(c *gin.Context) {
	name := c.Query("name")
	stock := c.Query("stock")
	SKU := c.Query("SKU")
	prize := c.Query("prize")

	var db, err = initializers.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("insert into products (`name`, `stock`, `SKU`, `prize`) values (?, ?, ?, ?)",
		name,
		stock,
		SKU,
		prize)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"stock": stock,
		"SKU":   SKU,
		"prize": prize,
	})

}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var db, err = initializers.Connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	_, err = db.Exec("delete from products where id=? ", id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}
