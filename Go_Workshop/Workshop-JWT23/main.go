package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/workshop-jwt2/controllers"
	"github.com/workshop-jwt2/database"
	"github.com/workshop-jwt2/middlewares"
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Println("please consider environment variable: %s", err)
	}

	// Initialize Database
	database.ConnectDatabase("test.db")
	database.Migrate()
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
