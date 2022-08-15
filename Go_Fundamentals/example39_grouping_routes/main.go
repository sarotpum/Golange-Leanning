package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Authen group
	// localhost:85/v1/authen/profile

	v1 := router.Group("v1")
	{
		authenAPI := v1.Group("/authen")
		{
			authenAPI.GET("/login", login)
			authenAPI.GET("/register", register)
			authenAPI.GET("/profile", func(c *gin.Context) {
				c.String(http.StatusOK, "Profile")
			})
		}
	}

	stockAPI := router.Group("/stock")
	{
		stockAPI.GET("/list", listProduct)
		stockAPI.GET("/create", createProduct)
	}

	router.Run(":85")
}

func login(c *gin.Context) {
	c.String(http.StatusOK, "Login")
}

func register(c *gin.Context) {
	c.String(http.StatusOK, "Register")
}

//----------- Stock
func listProduct(c *gin.Context) {
	c.String(http.StatusOK, "List Product")
}

func createProduct(c *gin.Context) {
	c.String(http.StatusOK, "Create Product")
}
