package main

import (
	"net/http"

	"vss/controllers"
	"vss/middleware"
	"vss/product"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/product/:id", product.GetSingleProduct)
	router.POST("/product/:id", product.UpdateProduct)
	router.POST("/deleteproduct/:id", product.DeleteProduct)
	router.POST("/newproduct", product.NewProduct)

	router.POST("/login", controllers.Login)
	private := router.Group("/api", middleware.RequireAuth)
	{
		private.GET("/val", secureRoot)
	}

	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": "index",
	})
}

func secureRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"secure": "login",
	})
}
