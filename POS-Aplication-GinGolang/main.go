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
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	router.GET("/prod/:id", product.GetSingleProduct)

	private := router.Group("/api", middleware.RequireAuth)
	{
		private.GET("/product/:id", product.GetSingleProduct)
		private.POST("/product/:id", product.UpdateProduct)
		private.POST("/deleteproduct/:id", product.DeleteProduct)
		private.POST("/newproduct", product.NewProduct)
	}

	router.Run(":" + os.Getenv("PORT"))
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"page": "index",
	})
}
